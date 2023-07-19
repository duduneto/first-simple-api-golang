package books_routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Book struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Title     string
	Genre     string
}

type ListBookResponse struct {
	Total int
	Data  []Book
}

var listOfBooks []Book = []Book{}

func books(writeResponse http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId, _ := vars["id"]
	fmt.Println(bookId)

	switch request.Method {
	case "POST":
		registerBooks(writeResponse, request)
	case "GET":
		listBooks(writeResponse, request)
	default:
		fmt.Fprint(writeResponse, "ROUTE NOT FOUND")
	}
}

func registerBooks(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	newBook.Id = uuid.New()
	newBook.CreatedAt = time.Now()
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	// Close the request body to prevent resource leaks
	defer r.Body.Close()

	// Append the new book to the listOfBooks
	listOfBooks = append(listOfBooks, newBook)

	// Send a response
	w.WriteHeader(http.StatusCreated)
}

func listBooks(writeResponse http.ResponseWriter, request *http.Request) {
	bufResponse, _errConvertToByte := json.Marshal(
		ListBookResponse{
			Total: len(listOfBooks),
			Data:  listOfBooks,
		},
	)
	if _errConvertToByte != nil {
		http.Error(writeResponse, "Failed to convert into JSON", http.StatusInternalServerError)
		return
	}
	writeResponse.Header().Set("Content-Type", "application/json")
	_, err := writeResponse.Write(bufResponse)
	if err != nil {
		http.Error(writeResponse, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}

func getSingle(writeResponse http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId, err := vars["id"]
	fmt.Println(bookId)
	if err != false {
		http.Error(writeResponse, "Pass a valid id", http.StatusInternalServerError)
		return
	}
	var found = Book{}
	for _, book := range listOfBooks {
		if book.Id.String() == bookId {
			found = book
		}
	}
	generateResponseRequest(found, writeResponse, request)
}

func generateResponseRequest(payload any, writeResponse http.ResponseWriter, request *http.Request) {
	bufResponse, _errConvertToByte := json.Marshal(payload)
	if _errConvertToByte != nil {
		http.Error(writeResponse, "Failed to convert into JSON", http.StatusInternalServerError)
		return
	}
	writeResponse.Header().Set("Content-Type", "application/json")
	_, err := writeResponse.Write(bufResponse)
	if err != nil {
		http.Error(writeResponse, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}

func HandleBooksRoutes(server *http.ServeMux) {
	server.HandleFunc("/books", books)
}
