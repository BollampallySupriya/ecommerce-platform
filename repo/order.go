package repo

import (
	"context"
	"time"
)

type Order struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Customer        uint64   `json:"customer_id"`
	Price           float64  `json:"price"`
	LineItems       []uint64 `json:"line_items"`
	DeliveryAddress string   `json:"delivery_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func (DB *DB) ListAllOrders(ctx context.Context) ([]*Order, error){

	query := `SELECT * from order;`

	rows, err := DB.Conn.Query(ctx, query)

	// rows, err := db.QueryContext(ctx, query) use with sql.DB 
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

func (DB *DB) PostOrder(ctx context.Context, order *Order) (*Order, error) {
	query := `
		INSERT INTO orders (id, name, customer_id, price, line_items, delivery_address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, customer_id, price, line_items, delivery_address, created_at, updated_at;
	`

	var newOrder Order
	err := DB.Conn.QueryRow(ctx, query,
		order.ID,
		order.Name,
		order.Customer,
		order.Price,
		order.LineItems,
		order.DeliveryAddress,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(
		&newOrder.ID,
		&newOrder.Name,
		&newOrder.Customer,
		&newOrder.Price,
		&newOrder.LineItems,
		&newOrder.DeliveryAddress,
		&newOrder.CreatedAt,
		&newOrder.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newOrder, nil
}


func (DB *DB) UpdateOrder(ctx context.Context, orderID string, updateOrder *Order) (*Order, error) {

	query := `UPDATE orders SET name = COALESCE($1, name), 
                                customer_id = COALESCE($2, customer_id), 
                                price = COALESCE($3, price), 
                                line_items = COALESCE($4, line_items), 
                                delivery_address = COALESCE($5, delivery_address), 
                                updated_at = NOW() 
              WHERE id = $6 RETURNING id, name, customer_id, price, line_items, delivery_address, updated_at;`

	var updatedOrder Order

	err := DB.Conn.QueryRow(ctx, query,
				updateOrder.Name,
				updateOrder.Customer,
				updateOrder.Price,
				updateOrder.LineItems,
				updateOrder.DeliveryAddress,
				orderID,
			).Scan(
				&updatedOrder.ID,
				&updatedOrder.Name,
				&updatedOrder.Customer,
				&updatedOrder.Price,
				&updatedOrder.LineItems,
				&updatedOrder.DeliveryAddress,
				&updatedOrder.UpdatedAt,
			)
	
	if err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}