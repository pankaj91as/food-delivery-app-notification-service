package repository

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
	"log"
)

func (r *Repository) GetOrders(ctx context.Context) ([]model.Orders, error) {
	var orders []model.Orders
	query := "select id, customer_id, order_status from `orders`"
	err := r.db.SelectContext(ctx, &orders, query)
	if err != nil {
		log.Println("Error while getting orders from database")
		return nil, err
	}
	return orders, nil
}

func (r *Repository) GetOrderByID(ctx context.Context, OrderID string) ([]model.Orders, error) {
	var orders []model.Orders
	query := "select id, customer_id, order_status from `orders` where id=?"
	err := r.db.SelectContext(ctx, &orders, query, OrderID)
	if err != nil {
		log.Println("Error while getting orders from database")
		return nil, err
	}
	return orders, nil
}

func (r *Repository) UpdateOrderByID(ctx context.Context, OrderID, OrderStatus string) (int64, error) {
	query := "UPDATE `orders` SET `order_status`=? WHERE `id`=?"
	result := r.db.MustExecContext(ctx, query, OrderStatus, OrderID)
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("Error while getting orders from database")
		return 0, err
	}
	return rows, nil
}
