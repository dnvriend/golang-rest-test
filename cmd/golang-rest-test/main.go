package main

import (
	"github.com/dnvriend/golang-rest-test/internal/dao"
	"github.com/dnvriend/golang-rest-test/internal/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	dao.LoadBooks()
	r.HandleFunc("/api/books", rest.GetBooks).Methods("GET")
	r.HandleFunc("/api/books", rest.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", rest.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", rest.DeleteBook).Methods("DELETE")
	r.HandleFunc("/api/books", rest.UpdateBook).Methods("PUT")
	log.Print("Running Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
