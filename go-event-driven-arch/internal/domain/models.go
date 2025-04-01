package domain

type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Products   []Product `json:"products"`
	Status     string    `json:"status"`
}

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
