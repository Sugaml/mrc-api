package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Image         string `json:"image"`
	IsAdmin       bool   `json:"is_admin"`
	Active        bool   `json:"active"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
}

func (user *User) Prepare() {
	user.ID = 0
	user.Active = true
	user.EmailVerified = true
	user.IsAdmin = false
}

func (user *User) Validate() error {
	if user.Email == "" {
		return errors.New("required email")
	}
	// if user.Username == "" {
	// 	return errors.New("required username")
	// }
	if user.Password == "" {
		return errors.New("required password")
	}
	return nil
}

func (user *User) LoginValidate() error {
	if user.Username == "" {
		return errors.New("required username")
	}
	if user.Password == "" {
		return errors.New("required password")
	}
	return nil
}
