package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirstName      string `gorm:"first_name" json:"first_name"`
	LastName       string `gorm:"last_name" json:"last_name"`
	Gender         string `gorm:"gender" json:"gender"`
	DOB            string `gorm:"dob" json:"dob"`
	MobileNumber   string `gorm:"mobile_num" json:"mobile_num"`
	Email          string `gorm:"email" json:"email"`
	ParanetName    string `gorm:"parent_name" json:"parent_name"`
	ParanetMobile  string `gorm:"parent_mobile" json:"parent_mobile"`
	ParentRelation string `gorm:"parent_relation" json:"parent_relation"`
	CID            uint   `gorm:"cid" json:"cid"`
	Course         Course ` json:"course"`
}
