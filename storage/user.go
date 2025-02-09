package storage

import (
	"user_wallet/db"
	structs "user_wallet/struct"
)

type UserStorage struct{}

func (us *UserStorage) GetUser(userId int) (user structs.User, err error) {

	err = db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (us *UserStorage) DeleteUser(userId int) error {
	if err := db.DB.Delete(&structs.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}
func (us *UserStorage) UpdateUser(user *structs.User) (err error) {

	if err := db.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
func (us *UserStorage) GetAllUsers() ([]*structs.User, error) {
	var users []*structs.User

	if err := db.DB.Order("id asc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (us *UserStorage) CreateUser(user *structs.User) error {
	return db.DB.Create(&user).Error
}

func (us *UserStorage) GetUserByEmail(email string) (structs.User, error) {

	user := structs.User{}
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
