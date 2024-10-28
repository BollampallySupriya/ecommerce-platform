package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/ecommerce-platform/helpers"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}


func ConnectDB(ctx context.Context ,cfg helpers.Config) (*DB, error) {
	fmt.Println(cfg.DATABASE_URL)
	config, err := pgx.ParseConfig(cfg.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(config)

	dbConn, err := pgx.ConnectConfig(ctx, config)

	if err != nil {
		log.Fatalf("Cannot connect to database!!! %v", err)
		return nil, err
	}

	err = testDB(dbConn, ctx)

	if err != nil {
		log.Fatal("Cannot ping to database!!!")
		return nil, err
	}

	return &DB{Conn: dbConn}, nil
}

func testDB(dbConn *pgx.Conn, ctx context.Context) error {
	err := dbConn.Ping(ctx)
	if err != nil {
		return err 
	}
	return nil
}