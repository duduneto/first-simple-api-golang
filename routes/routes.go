package routes

import (
	"net/http"

	books_routes "library.com/api/routes/books"
)

func HandleRoutes(server *http.ServeMux) {
	books_routes.HandleBooksRoutes(server)
}
