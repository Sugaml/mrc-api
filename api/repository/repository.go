package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func Migrate(r *Repository) {
	r.DB.AutoMigrate(
		models.Course{},
		models.User{},
		models.UserRole{},
		models.Student{},
		models.Address{},
		models.StudentEducation{},
		models.StudentFile{},
		models.ResetPassword{},
	)
}
