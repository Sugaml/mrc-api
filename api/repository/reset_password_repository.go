package repository

import (
	"html"
	"strings"
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
)

type ResetPasswordInterface interface {
	SaveDatails(db *gorm.DB, data *models.ResetPassword) (*models.ResetPassword, error)
	FindByToken(db *gorm.DB, token string) (*models.ResetPassword, error)
	DeleteDetails(db *gorm.DB, data *models.ResetPassword) (int64, error)
}
type ResetPasswordType struct {
}

func NewResetPassword() ResetPasswordInterface {
	return &ResetPasswordType{}
}

func Prepare(resetPassword *models.ResetPassword) {
	resetPassword.Token = html.EscapeString(strings.TrimSpace(resetPassword.Token))
	resetPassword.Email = html.EscapeString(strings.TrimSpace(resetPassword.Email))
}

func (d *ResetPasswordType) SaveDatails(db *gorm.DB, resetPassword *models.ResetPassword) (*models.ResetPassword, error) {
	err := db.Create(&resetPassword).Error
	if err != nil {
		return &models.ResetPassword{}, err
	}
	return resetPassword, nil
}

func (d *ResetPasswordType) FindByToken(db *gorm.DB, token string) (*models.ResetPassword, error) {
	resetPassword := models.ResetPassword{}
	err := db.Model(models.ResetPassword{}).Where("token = ?", token).Take(&resetPassword).Error
	if err != nil {
		return nil, err
	}
	return &resetPassword, nil
}

func (d *ResetPasswordType) DeleteDetails(db *gorm.DB, resetPassword *models.ResetPassword) (int64, error) {
	db = db.Model(&models.ResetPassword{}).Where("id = ?", resetPassword.ID).Take(&models.ResetPassword{}).Delete(&models.ResetPassword{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
