package repository

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
	"log"
)

func (r *Repository) GetCustomers(ctx context.Context) ([]model.Customers, error) {
	var customers []model.Customers
	query := "SELECT `id`, `name`, `mobile`, `email`, `status` FROM `customers`"
	err := r.db.SelectContext(ctx, &customers, query)
	if err != nil {
		log.Println("Error while getting customers from database")
		return nil, err
	}
	return customers, nil
}
