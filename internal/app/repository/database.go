package repository

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	GetOrders(ctx context.Context) ([]model.Orders, error)
	GetOrderByID(ctx context.Context, OrderID string) ([]model.Orders, error)
	UpdateOrderByID(ctx context.Context, OrderID, OrderStatus string) (int64, error)

	GetCustomers(ctx context.Context) ([]model.Customers, error)
	GetCustomerByID(ctx context.Context, customerId string) ([]model.Customers, error)

	InsertNotification(ctx context.Context, notification model.Notifications) error
	GetNotification(ctx context.Context, notification model.Notifications, notificationStatus string) ([]model.Notifications, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepoInit(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}
