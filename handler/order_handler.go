package handler

import (
	"fmt"
	"net/http"
	"github.com/ecommerce-platform/repository"
)

type Order struct {
	Repo *repository.OrderRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Order create method")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Update order method")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In List orders method")
}

func (o * Order) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In get order method")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In delete order method")
}
