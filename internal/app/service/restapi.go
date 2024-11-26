package service

import (
	"context"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/pkg/model"
)

type IRestService interface {
	GetOrders(ctx context.Context) []model.Orders
	GetOrdersByID(ctx context.Context, OrderID string) []model.Orders
	UpdateOrderByID(ctx context.Context, OrderID, OrderStatus string) (int64, error)

	GetCustomers(ctx context.Context) []model.Customers
}

type RestService struct {
	restRepo repository.IRepository
}

func NewRestService(restRepo repository.IRepository) IRestService {
	return &RestService{
		restRepo: restRepo,
	}
}
