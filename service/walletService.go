package service

import (
	"errors"
	"user_wallet/storage"
	structs "user_wallet/struct"
)

type WalletService struct {
	storage storage.WalletStorage
}

func (ws *WalletService) GetUserWallet(userId int) (structs.User, structs.Wallet, error) {

	var wallet structs.Wallet

	uService := UserService{}
	user, err := uService.FindUser(userId)
	if err != nil {
		return user, wallet, err
	}

	wStorage := storage.WStorage{}

	wallet, err = wStorage.GetWallet(uint(userId))

	if err != nil {
		return user, wallet, errors.New("user does not have a wallet")
	}

	return user, wallet, nil

}

func (ws *WalletService) CreateWallet(walletCreateRQ structs.WalletCreateRQ) (structs.Wallet, error) {
	var wallet structs.Wallet

	if walletCreateRQ.UserId != nil {
		wallet.UserId = *walletCreateRQ.UserId
	}
	if walletCreateRQ.Balance != nil {
		wallet.Balance = *walletCreateRQ.Balance
	}

	uSorage := storage.UStorage{}
	_, err := uSorage.GetUser(int(wallet.UserId))
	if err != nil {
		return wallet, err
	}

	wStorage := storage.WStorage{}

	_, err = wStorage.GetWallet(wallet.UserId)

	if err == nil {
		return wallet, errors.New("user wallet already exists")
	}

	if err := wStorage.CreateWallet(&wallet); err != nil {
		return wallet, err
	}

	return wallet, nil

}
