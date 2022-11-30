package postgres

import "github.com/jinzhu/gorm"

type PConnect = struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) *PConnect {
	return &PConnect{
		DB: db,
	}
}
