package controller

import (
	"context"
	"food-delivery-app-notification-service/internal/app/service"
	"food-delivery-app-notification-service/pkg/model"
)

type ISubscriberController interface {
	PrepairNotification(ctx context.Context, payload *model.MQPayload) string
	SaveNotification(ctx context.Context, payload *model.MQPayload, actualMessage string) error

	SendNotification(ctx context.Context, message, notificationType string) error
}

type SubscriberController struct {
	subService service.ISubscriberService
}

func NewSubscriberController(subService service.ISubscriberService) ISubscriberController {
	return &SubscriberController{
		subService: subService,
	}
}
