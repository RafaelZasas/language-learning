package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/meilisearch/meilisearch-go"
)

// Define the types for an article using a struct

type Movie struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Plot     string `json:"plot"`
	Rated    string `json:"rated"`
	Released string `json:"released"`
	Runtime  string `json:"runtime"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Actors   string `json:"actors"`
	Language string `json:"lanuage"`
	Country  string `json:"country"`
	Awards   string `json:"awards"`
	Poster   string `json:"poster"`
	Ratings  []struct {
		Source string `json:"source"`
		Value  string `json:"value"`
	}
	Metascore  string `json:"metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbId     string `json:"imdbId"`
	Type       string `json:"type"`
	DVD        string `json:"dvd"`
	BoxOffice  string `json:"boxOffice"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"response"`
}

// Declaring a global array of movies that will be populated in the main function to sim a db
var Movies []Movie

var dbClient *meilisearch.Client
var moviesIndex *meilisearch.Index

func importMovies() {
	moviesFile, err := os.Open("./static/movies.json")

	if err != nil {
		fmt.Printf("Error importing movies file\n%v\n", err.Error())
	}

	jsonParser := json.NewDecoder(moviesFile)

	if err = jsonParser.Decode(&Movies); err != nil {
		fmt.Printf("There was an error parsing the movies file\n%v\n", err.Error())
	}

}

func handleRequests() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/movies/", handleMovies)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// This function will be hit for /articles and /articles/{id}
func handleMovies(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movies/")

	// if all the movies are being requested
	if id == "" && r.Method == http.MethodGet {
		getAllMovies(w, r)
		return
	}
	// convert path to integer
	id_num, err := strconv.Atoi(id)

	// resolve with a 400 malformed input response if the path is not an int
	if err != nil {
		http.Error(w, "The requested article must be a number", http.StatusUnprocessableEntity)
		return
	}

	// Decrementing id here to maintain readability
	getMovie(w, id_num-1)

}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Movies) // this will encode the articles into a json format
}

func getMovie(w http.ResponseWriter, id int) {

	if id >= len(Movies) {
		http.Error(w, "The requested entity does not exist", http.StatusNotFound)
		return
	}

	// addMovieToMeilieSearch(Movies[id])

	json.NewEncoder(w).Encode(Movies[id])
}

func addMovieToMeilieSearch(movie Movie) {
	task, err := moviesIndex.AddDocuments(movie)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Adding Movie to Meilie Search\nTask ID: %v\nstatus: %v\n", task.TaskUID, task.Status)
}

func main() {

	// interfacing with the meiliesearch db
	dbClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})

	// creating an index for movies
	// if the index "Movies" does not exist, one will be created as soon as a doc is added
	moviesIndex = dbClient.Index("movies")

	importMovies()
	handleRequests()
}
