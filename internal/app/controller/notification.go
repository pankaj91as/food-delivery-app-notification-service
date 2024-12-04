package controller

import (
	"context"
	"fmt"
	"food-delivery-app-notification-service/pkg/model"
	"strings"
)

func (subCon *SubscriberController) PrepairNotification(ctx context.Context, payload *model.MQPayload) string {
	fmt.Println("⚡⚙️ Prepairing Notification")
	customer := subCon.subService.GetCustomerByID(ctx, payload.CustomerID)
	newMessage := strings.Replace(payload.Message, "{user}", customer[0].Name, -1)
	newMessage = strings.Replace(newMessage, "{orderstate}", payload.OrderStatus, -1)
	return newMessage
}

func (subCon *SubscriberController) SaveNotification(ctx context.Context, payload *model.MQPayload, actualMessage string) error {
	fmt.Println("Saving Notification In Database!")

	notification := model.Notifications{
		OrderID:                payload.OrderID,
		CustomerID:             payload.CustomerID,
		NotificationTemplateID: "1001",
		NotificationChannel:    payload.NotificationType,
		NotificationStatus:     "prepared",
	}

	existingNotification, err := subCon.subService.GetNotifications(ctx, notification, "prepared")
	if err != nil {
		return err
	}

	if len(existingNotification) == 0 {
		err = subCon.subService.InsertNotifications(ctx, notification)
		if err != nil {
			return err
		}
	}
	return nil
}
