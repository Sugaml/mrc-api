package models

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `json:"student"`
	CourseID  uint    `json:"course_id"`
	Course    Course  `json:"course"`
	Level     string  `json:"level"`
	Semester  string  `json:"semster"`
	IsAdmit   bool    `json:"is_admit"`
	Remarks   string  `json:"remarks"`
}
