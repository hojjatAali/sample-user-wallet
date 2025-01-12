package models

type Wallet struct {
	ID      uint    `json:"id"`
	UserId  uint    `json:"userId"`
	Balance float64 `json:"balance"`
}
