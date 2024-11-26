package service

import (
	"context"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/pkg/model"
)

type IRestService interface {
	GetOrders(ctx context.Context) []model.Orders
}

type RestService struct {
	restRepo repository.IRepository
}

func NewRestService(restRepo repository.IRepository) IRestService {
	return &RestService{
		restRepo: restRepo,
	}
}
