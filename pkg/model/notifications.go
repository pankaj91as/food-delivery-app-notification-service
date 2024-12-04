package model

type Notifications struct {
	OrderID                string `json:"order_id" db:"order_id"`
	CustomerID             string `json:"customer_id" db:"customer_id"`
	NotificationTemplateID string `json:"notification_template_id" db:"notification_template_id"`
	NotificationChannel    string `json:"notification_channel" db:"notification_channel"`
	NotificationStatus     string `json:"notification_status" db:"notification_status"`
}
