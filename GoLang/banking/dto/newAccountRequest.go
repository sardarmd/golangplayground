package dto

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
	AccountType string  `json:"account_type"`
}
