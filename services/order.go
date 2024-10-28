package services

import (
	"context"
	"time"
)

type Order struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Customer        uint64   `json:"customer"`
	Price           float64  `json:"price"`
	LineItems       []uint64 `json:"lineItems"`
	DeliveryAddress string   `json:"deliveryAddress"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func (o *Order) getAllOrders() ([]*Order, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * from order;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err 
	}
	var orders []*Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Name,
			&order.Customer,
			&order.Price,
			&order.LineItems,
			&order.DeliveryAddress,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err 
		}
		orders = append(orders, &order)
	}
	return orders, nil
}