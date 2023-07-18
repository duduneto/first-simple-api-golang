package books_routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func books(writeResponse http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		registerBooks(writeResponse, request)
	default:
		fmt.Fprint(writeResponse, "ROUTE NOT FOUND")
	}
}

func registerBooks(writeResponse http.ResponseWriter, request *http.Request) {
	body, err_read_body := ioutil.ReadAll(request.Body)
	if err_read_body != nil {
		http.Error(writeResponse, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	// Close the request body to prevent resource leaks
	defer request.Body.Close()
	_, err := writeResponse.Write(body)
	if err != nil {
		http.Error(writeResponse, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}

func listBooks() {

}

func HandleBooksRoutes(server *http.ServeMux) {
	server.HandleFunc("/books", books)
}
