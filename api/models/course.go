package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Name         string `gorm:"name" json:"name"`
	Discription  string `gorm:"discription" json:"discription"`
	Fee          uint   `gorm:"fee" json:"fee"`
	Duration     uint   `gorm:"duration" json:"duration"`
	CreditHours  uint   `gorm:"credit_hours" json:"credit_hours"`
	CourseType   string `gorm:"course_type" json:"course_type"`
	AffiliatedBy string `gorm:"affiliated_by" json:"affiliated_by"`
	Quota        int    `gorm:"quota" json:"quota"`
	IsActive     bool   `gorm:"is_active" json:"is_active"`
}

func (course *Course) Validate() error {
	if course.Name == "" {
		return errors.New("required course name")
	}
	if course.Fee <= 0 {
		return errors.New("fee greater than 0")
	}
	if course.Duration == 0 {
		return errors.New("duration required")
	}
	if course.CourseType == "" {
		return errors.New("required course type name")
	}

	return nil
}
