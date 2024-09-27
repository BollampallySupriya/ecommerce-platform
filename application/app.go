package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ecommerce-platform/router"
	"github.com/jackc/pgx/v5"
)

type App struct {
	router http.Handler
	config Config
	rdb    *pgx.Conn
}

func New(config Config) *App {
	app := &App{}
	app.config = config
	app.router = router.LoadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		a.config.DBUser, a.config.DBPassword, a.config.DBName, a.config.DBHost, a.config.DBPort, a.config.DBSSLMode)

	// Connect to the database
	var err error
	a.rdb, err = pgx.Connect(ctx, connStr)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	// Ensure the connection is closed on exit
	defer func() {
		if err := a.rdb.Close(ctx); err != nil {
			log.Fatal("Error closing database connection: %w", err)
		}
	}()

	server := &http.Server{
		Addr: a.config.ServerAddr,
		Handler: a.router,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to listen and serve %w", err)
	}
	return nil

}
