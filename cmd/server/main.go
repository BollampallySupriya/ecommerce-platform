package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ecommerce-platform/db"
	"github.com/ecommerce-platform/router"
	"github.com/ecommerce-platform/services"
)

type Config struct{
	Port string 
}

type Application struct {
	Config Config
	Models services.Models
}

var port = os.Getenv("PORT") // 8080

func (app *Application) Serve() error {
	
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return server.ListenAndServe()
}


func main() {
	var cfg Config 
	cfg.Port = port 

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		fmt.Println("Error connecting DB")
	}
	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		fmt.Println("Error starting server")
	}
}

// installations :

// go get github.com/jackc/pgconn
// go get github.com/jackc/pgx/v4
// go get github.com/jackc/pgx/v4/stdlib
// go get github.com/lib/pq