package service

import (
	"errors"
	"user_wallet/db"
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

	if err := db.DB.Create(&user).Error; err != nil {
		return user, errors.New("dont create user")
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

	if err := db.DB.Updates(&user).Error; err != nil {
		return user, errors.New("cant update user")
	}

	return user, err
}

func (uS *UserService) FindUser(userId int) (user structs.User, err error) {

	err = db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil

}
func (uS *UserService) GetUser(userId int) (structs.UserWalletResponse, error) {
	userWalletresponse := structs.UserWalletResponse{}

	user, err := uS.FindUser(userId)
	if err != nil {
		return userWalletresponse, errors.New("user not found")
	}
	userWalletresponse.User = user

	var wallet structs.Wallet
	if err := db.DB.First(&wallet, "user_id = ?", userId).Preload("User").Error; err != nil {
		return userWalletresponse, nil
	}

	userWalletresponse.Wallet = wallet

	return userWalletresponse, nil
}

func (uS *UserService) DeleteUser(userId int) error {
	_, err := uS.FindUser(userId)

	if err != nil {
		return errors.New("user not found")
	}

	if err := db.DB.Delete(&structs.User{}, userId).Error; err != nil {
		return errors.New("delete failed")
	}
	return nil
}

func (uS *UserService) GetUsers() (users []*structs.User, err error) {

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, errors.New("can not get users")
	}
	return users, nil

}
