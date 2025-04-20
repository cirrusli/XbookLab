package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username            string         `gorm:"size:50;not null;unique"`
	PasswordHash        string         `gorm:"size:255;not null"`
	Nickname            string         `gorm:"size:50"`
	Avatar              string         `gorm:"size:255"`
	Bio                 string         `gorm:"type:text"`
	Email               string         `gorm:"size:100;unique"`
	CategoryPreferences map[string]int `gorm:"type:json"`
	Following           []User         `gorm:"many2many:user_follows;joinForeignKey:user_id;joinReferences:follow_id"`
	Followers           []User         `gorm:"many2many:user_follows;joinForeignKey:follow_id;joinReferences:user_id"`
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

func FollowUser(userID uint, followID string) error {
	return DB.Model(&User{}).Where("id = ?", userID).Association("Following").Append(&User{Model: gorm.Model{ID: userID}})
}

func UnfollowUser(userID uint, followID string) error {
	return DB.Model(&User{}).Where("id = ?", userID).Association("Following").Delete(&User{Model: gorm.Model{ID: userID}})
}

func GetFollowing(userID uint) ([]User, error) {
	var user User
	if err := DB.Preload("Following").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Following, nil
}

func GetFollowers(userID uint) ([]User, error) {
	var user User
	if err := DB.Preload("Followers").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Followers, nil
}

func DeleteUser(id uint) error {
	return DB.Delete(&User{}, id).Error
}
