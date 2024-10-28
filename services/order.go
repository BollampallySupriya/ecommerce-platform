package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)



func (App *Application) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	all, err := App.Repo.ListAllOrders(context.Background(), )
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