package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ecommerce-platform/services"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var orders services.Order 
	all, err := orders.getAllOrders()
	if err != nil {
		fmt.Println("Error", err)
		return 
	}
	response, err :=json.Marshal(all)
	w.Write(response) 
}