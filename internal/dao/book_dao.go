package dao

import (
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func LoadBooks() {
	books = append(books, Book{ID: "1", Isbn: "abcd", Title: "title1", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "efgh", Title: "title2", Author: &Author{Firstname: "Steven", Lastname: "Larson"}})
	books = append(books, Book{ID: "3", Isbn: "ijkl", Title: "title3", Author: &Author{Firstname: "Bob", Lastname: "Stevenson"}})
}

func GetBook(id string) Book {
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return Book{}
}

func GetBooks() []Book {
	return books
}

func DeleteBook(id string) []Book {
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	return books
}

func CreateBook(book Book) Book {
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	return book
}

func AddBook(book Book) Book {
	books = append(books, book)
	return book
}

func UpdateBook(book Book) Book {
	books = DeleteBook(book.ID)
	return AddBook(book)
}
