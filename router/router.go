package router

import (
	"net/http"
	"context"
	"fmt"
	"time"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/ecommerce-platform/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/cors"
)

type Router struct {
	router http.Handler
    App *services.Application
}


func New(app *services.Application) *Router {
	router := &Router{
		App: app,
	}
    router.router = router.LoadRoutes()
	return router
}



func (router *Router) LoadRoutes() http.Handler {
	newRouter := chi.NewRouter()
	newRouter.Use(middleware.Logger)
	// newRouter.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"http://*", "https://*"},
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders: []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge: 300,
	// }))

	newRouter.Route("/api/v1/orders", router.loadOrderRoutes)

	return newRouter 
}

func (router *Router) loadOrderRoutes(orderRouter chi.Router) {
	orderRouter.Get("/", router.App.GetAllOrders)
	orderRouter.Post("/", router.App.CreateOrder)
	orderRouter.Put("/{id}", router.App.GetAllOrders)
	orderRouter.Delete("/{id}", router.App.GetAllOrders)
}

func (r *Router) Start(ctx context.Context, port string) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r.router,
	}

	// Creating a new context that can be canceled to manage shutdown
	stopCtx, stop := context.WithCancel(context.Background())
	defer stop()

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Capture interrupt signals to stop the server gracefully
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		select {
		case <-sigChan:
			log.Println("Received shutdown signal")
			stop() // Cancel the stopCtx
		case <-ctx.Done():
			stop() // Stop the server if parent ctx is canceled
		}
	}()

	// Wait for stopCtx to be canceled
	<-stopCtx.Done()
	log.Println("Shutting down server...")

	// Create a new context with a timeout for the server shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Shut down the server gracefully
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server Shutdown Failed: %v", err)
	}

	// Close the database connection after server shutdown
	if err := r.App.Repo.Conn.Close(ctx); err != nil {
		log.Printf("Failed to close DB: %v", err)
	}

	log.Println("Server and DB connections closed successfully.")
	return nil
}


	// ch := make(chan error, 1)
	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		ch <- fmt.Errorf("error occurred %w", err)
	// 	}
	// 	defer close(ch)
	// }()


	// select {
	// case err := <-ch:
	// 	return err
	// case <-ctx.Done():
	// 	timeout, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	// 	defer cancel()
	// 	return server.Shutdown(timeout)
	// }

