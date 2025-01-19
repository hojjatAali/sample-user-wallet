package service

import (
	"errors"
	"log"
	"user_wallet/storage"
	structs "user_wallet/struct"
)

type UserService struct{}

func (uS *UserService) CreateUser(userCR structs.UserCreateRQ) (structs.User, error) {
	var user structs.User
	if userCR.Name != "" {
		user.Name = userCR.Name
	}
	if userCR.Email != "" {
		user.Email = userCR.Email
	}

	uStorage := storage.UStorage{}

	err := uStorage.CreateUser(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}
func (uS *UserService) UpdateUser(userId int, userUpdateRQ structs.UserUpdateRQ) (structs.User, error) {
	user, err := uS.FindUser(userId)
	if err != nil {
		return user, errors.New("user not found")
	}

	if userUpdateRQ.Name != "" {
		user.Name = userUpdateRQ.Name
	}
	if userUpdateRQ.Email != "" {
		user.Email = userUpdateRQ.Email
	}

	uStorage := storage.UStorage{}

	err = uStorage.UpdateUser(&user)

	if err != nil {
		return user, errors.New("user update failed")
	}

	return user, err
}

func (uS *UserService) FindUser(userId int) (user structs.User, err error) {

	uStorage := storage.UStorage{}
	user, err = uStorage.GetUser(userId)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil

}
func (uS *UserService) GetUser(userId int) (structs.UserWalletResponse, error) {
	userWalletResponse := structs.UserWalletResponse{}

	user, err := uS.FindUser(userId)
	if err != nil {
		return userWalletResponse, err
	}
	userWalletResponse.User = user

	var wallet structs.Wallet

	wService := WalletService{}

	_, wallet, err = wService.GetUserWallet(userId)

	userWalletResponse.Wallet = wallet

	return userWalletResponse, nil
}

func (uS *UserService) DeleteUser(userId int) error {
	_, err := uS.FindUser(userId)

	if err != nil {
		return err
	}

	uStorage := storage.UStorage{}
	err = uStorage.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (uS *UserService) GetUsers() (users []*structs.User, err error) {

	uStorage := storage.UStorage{}
	users, err = uStorage.GetAllUsers()
	log.Print(users, err)
	if err != nil {
		return nil, err
	}

	return users, nil

}
