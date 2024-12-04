package service

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
	"log"
)

func (s *RestService) GetCustomers(ctx context.Context) []model.Customers {
	customers, err := s.restRepo.GetCustomers(ctx)
	if err != nil {
		log.Println(err)
	}
	return customers
}

func (s *SubscriberService) GetCustomerByID(ctx context.Context, customerId string) []model.Customers {
	customer, err := s.notificationRepo.GetCustomerByID(ctx, customerId)
	if err != nil {
		log.Println(err)
	}
	return customer
}
