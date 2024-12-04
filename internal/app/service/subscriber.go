package service

import (
	"context"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/pkg/model"
)

type ISubscriberService interface {
	GetNotifications(ctx context.Context, notification model.Notifications, notificationStatus string) ([]model.Notifications, error)
	InsertNotifications(ctx context.Context, notification model.Notifications) error
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
