package model

type Orders struct {
	OrderID     string `db:"id"`
	CustomerID  string `db:"customer_id"`
	OrderStatus string `db:"order_status"`
}
