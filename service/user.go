package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"user_wallet/storage"
	structs "user_wallet/struct"
)

type UserService struct {
	Storage storage.UserStorage
}

func (uS *UserService) CreateUser(rq structs.UserCreateRQ) (structs.User, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		return structs.User{}, err
	}

	user := structs.User{
		Name:     rq.Name,
		Email:    rq.Email,
		Password: string(hashPassword),
	}

	err = uS.Storage.CreateUser(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}
func (uS *UserService) UpdateUser(userId int, userUpdateRQ structs.UserUpdateRQ) (structs.User, error) {
	user, err := uS.FindUser(userId)
	if err != nil {
		return user, err
	}

	if userUpdateRQ.Name != "" {
		user.Name = userUpdateRQ.Name
	}
	if userUpdateRQ.Email != "" {
		user.Email = userUpdateRQ.Email
	}

	err = uS.Storage.UpdateUser(&user)

	if err != nil {
		return user, errors.New("user update failed")
	}

	return user, err
}

func (uS *UserService) FindUser(userId int) (user structs.User, err error) {

	user, err = uS.Storage.GetUser(userId)

	if err != nil {
		return user, err
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

	err = uS.Storage.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (uS *UserService) GetUsers() (users []*structs.User, err error) {

	users, err = uS.Storage.GetAllUsers()
	log.Print(users, err)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (uS UserService) Login(user structs.UserLoginRQ) (*structs.User, error) {

	getUser, err := uS.Storage.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}
	log.Println(user.Password, "kkkkkk", getUser.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password)); err != nil {

		return nil, err
	}

	return &getUser, nil
}
