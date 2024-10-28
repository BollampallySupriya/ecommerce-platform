package router

import (
	"net/http"
	"context"
	"fmt"
	"time"
	"log"
	"github.com/ecommerce-platform/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	router http.Handler
	App *services.Application
}

func New(app *services.Application) *Router {
	router := &Router{
		App: app,
	}
	router.LoadRoutes()
	return router
}



func (router Router) LoadRoutes() http.Handler {
	newRouter := chi.NewRouter()
	newRouter.Use(middleware.Recoverer)
	newRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	newRouter.Get("/api/v1/orders", router.App.GetAllOrders)

	return newRouter 
}

func (r *Router) Start(ctx context.Context, port string) error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: r.router,
	}

	go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %v", err)
        }
    }()

	defer func() {
		if err:= r.App.Repo.Conn.Close(ctx); err != nil {
			fmt.Println("Failed to close DB", err)
		}
	}()

    <-ctx.Done() // waits for ctx cancellation to gracefully shutdown
    log.Println("Shutting down server...")

    shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer shutdownCancel()
    return server.Shutdown(shutdownCtx)



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

}