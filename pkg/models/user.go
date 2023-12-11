package models

import (
	"github.com/mattot-the-builder/totgram-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID = ?", Id).Find(&user)
	return &user, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID = ?", ID).Delete(user)
	return user
}
