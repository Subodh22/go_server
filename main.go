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
	ID       string    `json:"ID"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
	return

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "34", Title: "Dune 2", Isbn: "ew22", Director: &Director{Firstname: "Subodh", Lastname: "Masha"}})
	movies = append(movies, Movie{ID: "43", Title: "Star wars", Isbn: "df", Director: &Director{Firstname: "Lalit", Lastname: "popo"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{ID}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting a server in port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
