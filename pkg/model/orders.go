package model

type Orders struct {
	OrderID     string `db:"id"`
	CustomerID  string `json:"customer_id,omitempty" db:"customer_id"`
	OrderStatus string `json:"order_status,omitempty" db:"order_status"`
}
