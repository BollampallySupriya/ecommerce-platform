package main

import (
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "os"
	// "os/signal"
	"github.com/go-chi/chi/v5"
)

type RouteResponse struct {
	Message string `json:"message"`
}

func main() {

	fmt.Println("Hello Riya!!!")

	router := *chi.NewRouter()

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(RouteResponse{Message: "Hello Riya!"})
	})
	log.Fatal(http.ListenAndServe(":5000", &router))
}

