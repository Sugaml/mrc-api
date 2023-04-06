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

type UserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `gorm:"gender" json:"gender"`
	Role      string `json:"role"`
	Image     string `json:"image"`
}

func NewUser(req *UserRequest) *User {
	return &User{
		Email:    req.Email,
		Username: req.Email,
		Password: req.Password,
		Role:     req.Role,
		Image:    req.Image,
	}
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
	if user.Password == "" {
		return errors.New("required password")
	}
	if len(user.Password) < 8 {
		return errors.New("password should be atleast 8 characters")
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
