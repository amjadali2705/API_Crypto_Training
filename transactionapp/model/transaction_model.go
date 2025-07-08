package model

type Transaction struct {
	ID            int     `json:"id"`
	FromAccountID int     `json:"from_account_id"`
	ToAccountID   int     `json:"to_account_id"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
}
