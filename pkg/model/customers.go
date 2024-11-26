package model

type Customers struct {
	CustomerID string `db:"id"`
	Name       string `db:"name"`
	Mobile     string `db:"mobile"`
	Email      string `db:"email"`
	Status     string `db:"status"`
}
