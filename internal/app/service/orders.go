package service

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
	"log"
)

func (s *RestService) GetOrders(ctx context.Context) []model.Orders {
	orders, err := s.restRepo.GetOrders(ctx)
	if err != nil {
		log.Println(err)
	}
	return orders
}
