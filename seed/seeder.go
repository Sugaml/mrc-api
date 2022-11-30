package seed

import (
	"log"

	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		FirstName:     "admin",
		LastName:      "admin",
		Email:         "admin@admin.com",
		Password:      "admin",
		IsAdmin:       true,
		Active:        true,
		EmailVerified: true,
	},
}

var roles = []models.UserRole{
	{
		Code:   1,
		Name:   "admin",
		Active: true,
	},
	{
		Code:   2,
		Name:   "write",
		Active: true,
	},
	{
		Code:   3,
		Name:   "read",
		Active: true,
	},
}

func Load(db *gorm.DB) {
	for _, user := range users {
		err := db.Debug().Model(&models.User{}).FirstOrCreate(&user, &models.User{
			Email: user.Email,
		}).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for _, role := range roles {
		err := db.Debug().Model(&models.UserRole{}).FirstOrCreate(&role, &models.UserRole{
			Code: role.Code,
		}).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
