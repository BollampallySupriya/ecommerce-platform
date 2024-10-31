package main

import (
	"context"
	"time"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ecommerce-platform/repo"
	"github.com/ecommerce-platform/helpers"
	"github.com/ecommerce-platform/services"
	"github.com/ecommerce-platform/router"
)




func main() {

	// time for db process with any transaction
	const dbTimeout = time.Second * 3

	// var cfg *helpers.Config 
	cfg := helpers.LoadConfig()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()


	dbConn, err := repo.ConnectDB(ctx, *cfg)

	if err != nil {
		log.Fatalf("Error While Connecting DB: %v", err)
	}
	app := services.New(dbConn)

	server := router.New(app)

	log.Printf("Starting server on port %s...", cfg.Port)
	server.Start(ctx, cfg.Port)
}