package domain

type Expense struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	PaymentType string  `json:"payment_type"`
	CreatedAt   string  `json:"created_at"`
}
