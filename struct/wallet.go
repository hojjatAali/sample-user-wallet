package structs

type Wallet struct {
	ID      uint    `json:"id"`
	UserId  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

type WallerCreateRQ struct {
	UserId  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}
