package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:50;not null;unique"`
	PasswordHash string `gorm:"size:255;not null"`
	Avatar       string `gorm:"size:255"`
	Email        string `gorm:"size:100;unique"`
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func GetUserByID(id uint) (User, error) {
	var user User
	err := DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return DB.Delete(&User{}, id).Error
}
