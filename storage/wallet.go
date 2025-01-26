package storage

import (
	"user_wallet/db"
	structs "user_wallet/struct"
)

type WalletStorage interface {
	CreateWallet(wallet *structs.Wallet) error
	GetWallet(userId uint) (wallet structs.Wallet, err error)
}
type WStorage struct{}

func (ws *WStorage) CreateWallet(wallet *structs.Wallet) error {

	return db.DB.Create(wallet).Error
}
func (ws *WStorage) GetWallet(userId uint) (wallet structs.Wallet, err error) {

	err = db.DB.Where("user_id = ?", userId).First(&wallet).Error

	return wallet, err
}
