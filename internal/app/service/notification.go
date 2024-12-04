package service

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
)

func (sub *SubscriberService) GetNotifications(ctx context.Context, notification model.Notifications, notificationStatus string) ([]model.Notifications, error) {
	notifications, err := sub.notificationRepo.GetNotification(ctx, notification, notificationStatus)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (sub *SubscriberService) InsertNotifications(ctx context.Context, notification model.Notifications) error {
	err := sub.notificationRepo.InsertNotification(ctx, notification)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriberService) GetNotificationByOrderCustomerAndTypeID(customerID, notificationType, orderID string) *model.Notifications {
	var notifications *model.Notifications
	return notifications
}
