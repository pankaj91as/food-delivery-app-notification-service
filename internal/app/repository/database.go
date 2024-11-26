package repository

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	GetOrders(ctx context.Context) ([]model.Orders, error)

	GetCustomers(ctx context.Context) ([]model.Customers, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepoInit(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}
