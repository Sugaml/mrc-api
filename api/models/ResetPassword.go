package models

import (
	"github.com/jinzhu/gorm"
)

type ResetPassword struct {
	gorm.Model
	Email string `gorm:"size:100;not null;" json:"email"`
	Token string `gorm:"size:255;not null;" json:"token"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}
