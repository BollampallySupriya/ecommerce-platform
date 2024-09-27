package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/ecommerce-platform/handler"
	// "github.com/ecommerce-platform/application"
)

func LoadRoutes() *chi.Mux{
	router := chi.NewRouter()

	router.Get("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("GOTO : \n /orders for Orders"))
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", handleOrderRoutes)

	return router
}

func handleOrderRoutes(router chi.Router) {
	// define handler here which handles all the crud operations.
}