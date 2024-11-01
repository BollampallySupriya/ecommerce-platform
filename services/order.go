package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ecommerce-platform/repo"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)



func (App *Application) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	all, err := App.Repo.ListAllOrders(context.Background())
	if err != nil {
		fmt.Println("Error", err)
		return 
	}
	response, err :=json.Marshal(all)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	w.Write(response) 
}

func (App *Application) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o struct {
		Name            string   `json:"name"`
		Customer        uint64   `json:"customer"`
		Price           float64  `json:"price"`
		LineItems       []uint64 `json:"lineItems"`
		DeliveryAddress string   `json:"deliveryAddress"`
	}
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		w.Write([]byte("Please check and provide correct data!!!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	now := time.Now().UTC()
	order := repo.Order{
		ID: uuid.New().String(),
		Name: o.Name,
		Customer: o.Customer,
		LineItems: o.LineItems,
		DeliveryAddress: o.DeliveryAddress,
		CreatedAt: now,
		UpdatedAt: now,
	}
	newOrder, err := App.Repo.PostOrder(context.Background(), &order)
	if err != nil {
		fmt.Println("Error", err)
		return 
	}
	response, err :=json.Marshal(newOrder)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	w.Write(response) 
}

func (App *Application) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var orderID = chi.URLParam(r, "id")

	var updateOrder *repo.Order

	if err := json.NewDecoder(r.Body).Decode(&updateOrder); err != nil {
		http.Error(w, "Error while updating order.", http.StatusBadRequest)
		return 
	}

	order, err := App.Repo.UpdateOrder(context.Background(), orderID, updateOrder)
	
	if err != nil {
		http.Error(w, "Error while updating order.", http.StatusBadRequest)
		return 
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (App *Application) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var orderID = chi.URLParam(r, "id")

	err := App.Repo.DeleteOrder(context.Background(), orderID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Deleted Order Successfully"))
}