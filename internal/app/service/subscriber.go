package service

import (
	"context"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/pkg/model"
)

type ISubscriberService interface {
	GetNotifications() *model.Notifications
	InsertNotifications() error
	GetNotificationByOrderCustomerAndTypeID(customerID, notificationType, orderID string) *model.Notifications

	GetCustomerByID(ctx context.Context, customerId string) []model.Customers
}

type SubscriberService struct {
	notificationRepo repository.IRepository
}

func NewSubscriberService(notificationRepo repository.IRepository) ISubscriberService {
	return &SubscriberService{
		notificationRepo: notificationRepo,
	}
}
