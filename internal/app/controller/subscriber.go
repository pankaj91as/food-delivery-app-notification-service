package controller

import (
	"context"
	"food-delivery-app-notification-service/internal/app/service"
	"food-delivery-app-notification-service/pkg/model"
)

type ISubscriberController interface {
	PrepairNotification(ctx context.Context, payload *model.MQPayload) string
}

type SubscriberController struct {
	subService service.ISubscriberService
}

func NewSubscriberController(subService service.ISubscriberService) ISubscriberController {
	return &SubscriberController{
		subService: subService,
	}
}
