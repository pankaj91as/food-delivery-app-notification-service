package service

import "food-delivery-app-notification-service/pkg/model"

func (sub *SubscriberService) GetNotifications() *model.Notifications {
	var notifications *model.Notifications
	return notifications
}

func (sub *SubscriberService) InsertNotifications() error {
	return nil
}

func (sub *SubscriberService) GetNotificationByOrderCustomerAndTypeID(customerID, notificationType, orderID string) *model.Notifications {
	var notifications *model.Notifications
	return notifications
}
