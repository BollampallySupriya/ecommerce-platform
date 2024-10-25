package application

import (
	"net/http"

	"github.com/ecommerce-platform/handler"
	"github.com/ecommerce-platform/repository"
	"github.com/go-chi/chi/v5"

	// "github.com/ecommerce-platform/application"
	"fmt"
	// "time"

	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	// "github.com/go-chi/oauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func (a *App) LoadRoutes(){
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false, // allows credentials if true and does not allow credentials if false.
		MaxAge:           300,   // Maximum value not ignored by any of major browsers
	}))

	// registerAPI(router)

	// Seek, verify and validate JWT tokens
	router.Use(jwtauth.Verifier(tokenAuth))

	// Handle valid / invalid tokens. In this example, we use
	// the provided authenticator middleware, but you can write your
	// own very easily, look at the Authenticator method in jwtauth.go
	// and tweak it, its not scary.
	// router.Use(jwtauth.Authenticator)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		w.Write([]byte("GOTO : \n /orders for Orders"))
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.handleOrderRoutes)

	a.router = router
}

// func registerAPI(r *chi.Mux) {
// 	s := oauth.NewBearerServer(
// 		"mySecretKey-10101",
// 		time.Second*120,
// 		&TestUserVerifier{},
// 		nil)
// 	r.Post("/token", s.UserCredentials)
// 	r.Post("/auth", s.ClientCredentials)
// }

func (a *App) handleOrderRoutes(router chi.Router) {
	// define handler here which handles all the crud operations.
	orderHandler := &handler.Order{
		Repo : &repository.OrderRepo{
			Client: a.rdb,
		},
	}
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.Get)
	router.Post("/", orderHandler.Create)
	router.Put("/{id}", orderHandler.Update)
	router.Delete("/{id}", orderHandler.Delete)
}
