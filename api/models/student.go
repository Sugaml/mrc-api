package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	FirstName      string `gorm:"first_name" json:"firstname"`
	LastName       string `gorm:"last_name" json:"lastname"`
	Gender         string `gorm:"gender" json:"gender"`
	DOB            string `gorm:"dob" json:"dob"`
	MobileNumber   string `gorm:"mobile_num" json:"mobile_num"`
	Email          string `gorm:"email" json:"email"`
	ParanetName    string `gorm:"parent_name" json:"parent_name"`
	ParanetMobile  string `gorm:"parent_mobile" json:"parent_mobile"`
	ParentRelation string `gorm:"parent_relation" json:"parent_relation"`
	CID            uint   `gorm:"cid" json:"cid"`
	Course         Course `gorm:"foreignkey:CID" json:"course"`
	UserId         uint   `gorm:"user_id" json:"user_id"`
	User           User   `gorm:"foreignkey:UserId" json:"user"`
	IsApproved     bool   `gorm:"is_approved;default:false" json:"is_approved"`
}

type StudentStatusRequest struct {
	Status bool `json:"status"`
}

func NewStudent(id uint, req *UserRequest) *Student {
	return &Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Gender:    req.Gender,
		UserId:    id,
	}
}

func (user *Student) Validate() error {
	if user.FirstName == "" {
		return errors.New("required first name")
	}
	if user.LastName == "" {
		return errors.New("required last name")
	}
	return nil
}
