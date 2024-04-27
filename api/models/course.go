package models

import (
	"errors"
	"time"

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
	Subject      uint   `gorm:"subject" json:"subject"`
	Faculty      string `gorm:"faculty" json:"faculty"`
	Year         uint   `gorm:"year" json:"year"`
	IsActive     bool   `gorm:"is_active" json:"is_active"`
}

type CourseResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	Name         string    `json:"name"`
	Discription  string    `json:"discription"`
	Fee          uint      `json:"fee"`
	Duration     uint      `json:"duration"`
	CreditHours  uint      `json:"creditHours"`
	CourseType   string    `json:"courseType"`
	AffiliatedBy string    `json:"affiliatedBy"`
	Quota        int       `json:"quota"`
	Subject      uint      `json:"subject"`
	Faculty      string    `json:"faculty"`
	Year         uint      `json:"year"`
	IsActive     bool      `json:"isActive"`
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
