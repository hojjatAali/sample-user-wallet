package structs

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Wallet struct {
	ID      uint    `json:"id"`
	UserId  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

type WalletCreateRQ struct {
	UserId  *uint    `json:"user_id" validate:"required"`
	Balance *float64 `json:"balance" validate:"required"`
}

type WalletUpdateRQ struct {
	UserId  *uint    `json:"user_id" validate:"required"`
	Balance *float64 `json:"balance" validate:"required"`
}
type UserWalletResponse struct {
	User   User
	Wallet Wallet
}

func (w *WalletCreateRQ) Validate() error {
	return validate.Struct(w)
}
func (w *WalletUpdateRQ) Validate() error {
	return validate.Struct(w)
}
