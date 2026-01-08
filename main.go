package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//header() is a map (x,y) -> content-type:application/json
	// its for the response body - > type:json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // It extracts path parameters from the URL.
	//map[string]string{
	//"id": "3",

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // a header is a piece of metadata, structured as a key-value pair, that provides additional context about the request or response message
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                            // struct of type Movie to get all the data as request to create a new movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //Decode(&movie)- sends the decoded json from the request body to the go struct
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

/*POST request
   ↓
Decode JSON → Go struct
   ↓
Validate
   ↓
Create resource
   ↓
Return created resource (201)*/

func updateMovie(w http.ResponseWriter, r *http.Request) {

	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over movies
	for index, item := range movies {
		if item.ID == params["id"] {
			//delete movie by id
			movies = append(movies[:index], movies[index+1:]...)
			//add a new movie sent in the body of postman
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "756453", Title: "xyz", Director: &Director{FirstName: "a", LastName: "b"}})
	movies = append(movies, Movie{ID: "2", Isbn: "123456", Title: "gfd", Director: &Director{FirstName: "v", LastName: "n"}})
	movies = append(movies, Movie{ID: "3", Isbn: "456789", Title: "tre", Director: &Director{FirstName: "q", LastName: "r"}})
	movies = append(movies, Movie{ID: "4", Isbn: "901234", Title: "qrq", Director: &Director{FirstName: "s", LastName: "t"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting the server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r)) // its stops the program in the case of any error
	// listenandserver is a function to start the server ( infinite loop to serve the requests)
	//it returns an error (if any)

	/*Client Request
	     ↓
	  http.ListenAndServe
	     ↓
	  Router (r)
	     ↓
	  Handler
	     ↓
	  Service
	     ↓
	  DB */

	//Router is a handler that calls other handlers.

}

// imp - If a function needs to change something, pass a pointer.
//Go passes values, not references.
//To modify something, pass its address.
