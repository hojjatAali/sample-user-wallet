package storage

import (
	"user_wallet/db"
	structs "user_wallet/struct"
)

type UserStorage interface {
	CreateUser(user *structs.User) error
	GetAllUsers() ([]*structs.User, error)
	UpdateUser(user *structs.User) (err error)
	DeleteUser(userId int) error
	GetUser(userId int) (user structs.User, err error)
}
type UStorage struct{}

func (us *UStorage) GetUser(userId int) (user structs.User, err error) {

	err = db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (us *UStorage) DeleteUser(userId int) error {
	if err := db.DB.Delete(&structs.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}
func (us *UStorage) UpdateUser(user *structs.User) (err error) {

	if err := db.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
func (us *UStorage) GetAllUsers() ([]*structs.User, error) {
	var users []*structs.User

	if err := db.DB.Order("id asc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (us *UStorage) CreateUser(user *structs.User) error {
	return db.DB.Create(&user).Error
}
