let x = 0;

const getMoviesButton = document.getElementById("get-movies-btn");
const moviesGrid = document.getElementById("movies-grid")

function incrementX() {
  x += 1;
  document.getElementById("count").innerText = x;
}

function decrementX() {
  x -= 1;
  document.getElementById("count").innerText = x;
  console.log("decremented");
}

let movies


function getMoviess() {
  fetch("http://localhost:8080/movies").then(async res => {
    movies = await res.json();
    htmlString = `<div class='block p-2 m-4'>${movies[0].title}</div>`;
    document.getElementById("get-movies-btn").hidden = true;
    document.getElementById("movies-grid").innerHTML = htmlString;
  }

  )

}

window.onload = (evt) => {
  fetch("http://localhost:8080/movies").then(async res => {
    movies = await res.json();
    htmlString = ""
    for (let movie of movies) {
      htmlString += `
    <div class='
    block py-4 px-6 m-4 bg-slate-800 text-slate-200 text-center
     align-middle rounded-md'>
     <div class='flex flex-col'>
     <p class='py-2'>
    ${movie.title}
    </p>
    <img src=${movie.poster}/>

    </div>

    </div>`;
    }
    document.getElementById("movies-grid").innerHTML = htmlString;
  })
}