package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
