package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/ecommerce-platform/helpers"
)

type DB struct {
	Conn *pgx.Conn
}


func ConnectDB(ctx context.Context ,cfg helpers.Config) (*DB, error) {
	config, err := pgx.ParseConfig(cfg.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dbConn, err := pgx.ConnectConfig(ctx, config)

	if err != nil {
		log.Fatal("Cannot connect to database!!!")
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