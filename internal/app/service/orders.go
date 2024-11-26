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

func (s *RestService) GetOrdersByID(ctx context.Context, OrderID string) []model.Orders {
	orders, err := s.restRepo.GetOrderByID(ctx, OrderID)
	if err != nil {
		log.Println(err)
	}
	return orders
}

func (s *RestService) UpdateOrderByID(ctx context.Context, OrderID, OrderStatus string) (int64, error) {
	orders, err := s.restRepo.UpdateOrderByID(ctx, OrderID, OrderStatus)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return orders, nil
}
