package model

type SampleData struct {
	Data string
}

type Response struct {
	Status  int
	Data    []interface{}
	Message string
}

type MQPayload struct {
	OrderID          string
	CustomerID       string
	Message          string
	NotificationType string
	QueueName        string
}
