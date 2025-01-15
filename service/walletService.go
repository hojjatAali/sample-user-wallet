package service

import (
	"errors"
	"user_wallet/db"
	structs "user_wallet/struct"
)

type WalletService struct{}

func (ws *WalletService) GetUserWallet(userId int) (structs.User, structs.Wallet, error) {

	var user structs.User
	var wallet structs.Wallet

	err := db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, wallet, errors.New("user not found")
	}

	err = db.DB.Where("user_id = ?", userId).First(&wallet).Error
	if err != nil {
		return user, wallet, errors.New("user does not have a wallet")
	}

	return user, wallet, nil

}

func (ws *WalletService) CreateWallet(walletCreateRQ structs.WalletCreateRQ) (structs.Wallet, error) {
	var wallet structs.Wallet
	var user structs.User

	if walletCreateRQ.UserId != nil {
		wallet.UserId = *walletCreateRQ.UserId
	}
	if walletCreateRQ.Balance != nil {
		wallet.Balance = *walletCreateRQ.Balance
	}

	err := db.DB.Where("id = ?", wallet.UserId).First(&user).Error
	if err != nil {
		return wallet, errors.New("user not found")
	}
	err = db.DB.Where("user_id = ?", wallet.UserId).First(&wallet).Error
	if err == nil {
		return wallet, errors.New("user wallet already exists")
	}

	if err := db.DB.Create(&wallet).Error; err != nil {
		return wallet, err
	}

	return wallet, nil

}
