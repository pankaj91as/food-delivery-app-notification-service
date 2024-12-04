package controller

import (
	"context"
	"fmt"
	"food-delivery-app-notification-service/pkg/model"
)

func (subCon *SubscriberController) PrepairNotification(ctx context.Context, payload *model.MQPayload) string {
	fmt.Println("Prepairing Notification")
	customer := subCon.subService.GetCustomerByID(ctx, payload.CustomerID)
	fmt.Println(customer)
	return ""
}
