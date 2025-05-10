package models

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:50;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	Avatar    string    `gorm:"size:255"`
	Bio       string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}
