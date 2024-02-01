package models

import (
	"sample/db"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserId    string    `json:"userId" mysql:"user_id"`
	UserName  string    `json:"userName" mysql:"user_name"`
	Age       uint8     `json:"age" mysql:"age"`
	Gender    string    `json:"gender" mysql:"gender"`
	UpdatedAt time.Time `json:"updatedAt" mysql:"updated_at"`
	CreatedAt time.Time `json:"createdAt" mysql:"created_at"`
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	if db.DB.Find(&users).Error != nil {
		return nil, echo.ErrNotFound
	}
	return users, nil
}

func GetUserById(id string) (*User, error) {
	user := User{}
	if db.DB.Where("user_id = ?", id).First(&user).Error != nil {
		return nil, echo.ErrNotFound
	}
	return &user, nil
}
