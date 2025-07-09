package model

type Transaction struct {
	ID            int     `json:"id"`
	FromAccountID int     `json:"from_account"`
	ToAccountID   int     `json:"to_account"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"createdAt"`
}
