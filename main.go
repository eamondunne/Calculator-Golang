package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/eamondunne/Calculator-Golang/pkg/add"
	"github.com/eamondunne/Calculator-Golang/pkg/subtract"
	"github.com/eamondunne/Calculator-Golang/pkg/multiply"
	"github.com/eamondunne/Calculator-Golang/pkg/divide"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
		fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	s := r.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("/add", add.AddHandler)
	s.HandleFunc("/subtract", subtract.SubtractHandler)
	s.HandleFunc("/multiply", multiply.MultiplyHandler)
	s.HandleFunc("/divide", divide.DivideHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
    handleRequests()
}