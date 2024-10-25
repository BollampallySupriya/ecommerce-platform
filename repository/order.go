package repository


import "github.com/jackc/pgx/v5"

type OrderRepo struct {
	Client *pgx.Conn
}


