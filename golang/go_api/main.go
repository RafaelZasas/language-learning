package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define the types for an article using a struct

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json":desc`
	Content string `json:"content`
}

// Declaring a global array of articles that will be populated in the main function to sim a db
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: Homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // this will encode the articles into a json format
}

func AldoPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/view", AldoPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	// populating the articles

	Articles = []Article{
		{
			Title:   "Golang",
			Desc:    "THis article is all about the Go Language",
			Content: "Go is a programming language made by Google and is fast."},
		{
			Title:   "Dart",
			Desc:    "This article is all about the Dart language",
			Content: "Dart is a programming language made by Google and is the foundation of the Flutter Framework"},
		{
			Title:   "Carbon",
			Desc:    "This article is all about the Carbon programming language",
			Content: "Carbon is yet another language made by Google and is supposed to be a successor to the might C++",
		},
	}
	handleRequests()
}
