package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type UserRole struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	Code        uint64 `gorm:"not null;" json:"code"`
	Description string `gorm:"size:1024;null" json:"description"`
	Active      bool   `gorm:"not null;" json:"active"`
}

func (data *UserRole) Prepare() {
	data.ID = 0
	data.Name = html.EscapeString(strings.TrimSpace(data.Name))
	data.Description = html.EscapeString(strings.TrimSpace(data.Description))
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
}

func (data *UserRole) Validate() error {
	if data.Name == "" {
		return errors.New("required Name")
	}
	if data.Description == "" {
		return errors.New("required Description")
	}
	if data.Code == 0 {
		return errors.New("required Code")
	}
	return nil
}
