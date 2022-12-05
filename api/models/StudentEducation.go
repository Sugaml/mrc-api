package models

import "github.com/jinzhu/gorm"

type StudentEducation struct {
	gorm.Model
	StudentID        uint   `json:"student_id"`
	InstituteName    string `json:"institute_name"`
	InstituteAddress string `json:"institute_address"`
	CourseName       string `json:"course_name"`
	Level            string `json:"level"`
	Grade            string `json:"grade"`
	GPA              string `json:"gpa"`
	CompletedYear    string `json:"completed_year"`
	Remarks          string `json:"remarks"`
}
