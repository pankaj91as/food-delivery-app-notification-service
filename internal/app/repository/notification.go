package repository

import (
	"context"
	"food-delivery-app-notification-service/pkg/model"
	"log"
)

func (r *Repository) InsertNotification(ctx context.Context, notification model.Notifications) error {
	query := "INSERT INTO `notifications`(`order_id`, `customer_id`, `notification_template_id`, `notification_channel`, `notification_status`) VALUES (?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query,
		notification.OrderID,
		notification.CustomerID,
		notification.NotificationTemplateID,
		notification.NotificationChannel,
		notification.NotificationStatus,
	)
	if err != nil {
		log.Fatalln("Error while saving notification in table")
		return err
	}
	return nil
}

func (r *Repository) GetNotification(ctx context.Context, notification model.Notifications, notificationStatus string) ([]model.Notifications, error) {
	var notifications []model.Notifications
	query := "SELECT `order_id`, `customer_id`, `notification_template_id`, `notification_channel`, `notification_status` FROM `notifications` where `order_id`=? and `customer_id`=? and `notification_status`=?"
	err := r.db.SelectContext(ctx, &notifications, query,
		notification.OrderID,
		notification.CustomerID,
		notificationStatus,
	)
	if err != nil {
		log.Println("Error while getting notifications from database", err)
		return nil, err
	}
	return notifications, nil
}
