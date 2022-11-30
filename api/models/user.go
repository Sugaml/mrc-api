package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	MobileNumber  int64  `json:"mobile_number"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	IsAdmin       bool   `json:"is_admin"`
	Active        bool   `json:"active"`
	EmailVerified bool   `json:"email_verified"`
	IsStudent     bool   `json:"is_student"`
}

func (user *User) Prepare() {
	user.ID = 0
	user.Active = true
	user.IsAdmin = false
}

func (user *User) Validate() error {
	if user.FirstName == "" {
		return errors.New("required firstname")
	}
	if user.LastName == "" {
		return errors.New("required lastname")
	}
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
