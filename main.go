package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Game struct {
	Position int    `json: "Position"`
	Name     string `json: "Name"`
	Score    int    `json: "Score"`
	Time     int    `json: "Time"`
}

type Games []Game

func allData(w http.ResponseWriter, r *http.Request) {
	game := Games{
		Game{Position: 0, Name: "Test position", Score: 0, Time: 0},
	}
	fmt.Println("Endpoint hit")
	json.NewEncoder(w).Encode(game)
}

func postAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test post")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/game", allData).Methods("GET")
	myRouter.HandleFunc("/game", postAllData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	handleRequests()
}
