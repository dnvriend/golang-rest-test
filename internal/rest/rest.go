package rest

import (
	"encoding/json"
	"github.com/dnvriend/golang-rest-test/internal/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func respond(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func getParam(r *http.Request, param string) string {
	params := mux.Vars(r)
	return params[param]
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	respond(w, dao.GetBooks())
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := getParam(r, "id")
	book := dao.GetBook(id)
	if book.ID == "" {
		w.WriteHeader(404)
	} else {
		respond(w, book)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := getParam(r, "id")
	respond(w, dao.DeleteBook(id))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book dao.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	respond(w, dao.CreateBook(book))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book dao.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	respond(w, dao.UpdateBook(book))
}
