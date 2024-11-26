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
