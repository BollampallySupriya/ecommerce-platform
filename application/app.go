package application

import (
	"context"
	"fmt"
	"time"
	"log"
	"net/http"
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
	app.LoadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", a.config.ServerAddr),
		Handler: a.router,
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		a.config.DBUser, a.config.DBPassword, a.config.DBName, a.config.DBHost, a.config.DBPort, a.config.DBSSLMode)

	// Connect to the database
	var err error
	a.rdb, err = pgx.Connect(ctx, connStr)
	// err = a.rdb.Ping(ctx)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	} else {
		fmt.Println("Started server !!!")
	}

	// Ensure the connection is closed on exit
	defer func() {
		if err := a.rdb.Close(ctx); err != nil {
			log.Fatal("Error closing database connection: %w", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("error occurred %w", err)
		}
		defer close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second * 10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
