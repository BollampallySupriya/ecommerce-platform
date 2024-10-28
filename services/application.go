package services

import (
	// "database/sql"
	"database/sql"
	"time"
)

type Application struct {
}

// time for db process with any transaction

const dbTimeout = time.Second * 3

var db *sql.DB

// TODO check if needed

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models {}
}

type Models struct {
	Order Order 
	JsonResponse JsonResponse
}