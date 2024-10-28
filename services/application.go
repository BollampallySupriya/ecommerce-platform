package services

import (
	"context"
	"github.com/ecommerce-platform/repo"
)

type App interface {
	ListAllOrders(ctx context.Context) ([]repo.Order, error)
}


type Application struct {
	Repo *repo.DB
}


func New(repo *repo.DB) *Application {
	return &Application{
		Repo: repo,
	}
}