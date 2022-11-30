package repository

import (
	"errors"

	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
)

type UserRoleInterface interface {
	Save(data *models.UserRole) (*models.UserRole, error)
	FindAll() (*[]models.UserRole, error)
	Find(pid uint64) (*models.UserRole, error)
	Update(data *models.UserRole) (*models.UserRole, error)
	Delete(uid uint64) (int64, error)
}

func NewUserRole() UserRoleInterface {
	return &Repository{}
}
func (r *Repository) Save(data *models.UserRole) (*models.UserRole, error) {
	err := r.DB.Model(&models.UserRole{}).Create(&data).Error
	if err != nil {
		return &models.UserRole{}, err
	}
	return data, nil
}

func (r *Repository) FindAll() (*[]models.UserRole, error) {
	datas := []models.UserRole{}
	err := r.DB.Model(&models.UserRole{}).Order("id").Find(&datas).Error
	if err != nil {
		return &[]models.UserRole{}, err
	}
	return &datas, nil
}

func (r *Repository) Find(pid uint64) (*models.UserRole, error) {
	data := models.UserRole{}
	err := r.DB.Model(&models.UserRole{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &models.UserRole{}, err
	}
	return &data, nil
}

func (r *Repository) Update(data *models.UserRole) (*models.UserRole, error) {
	err := r.DB.Model(&models.UserRole{}).Where("id = ?", data.ID).Updates(models.UserRole{
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
		Active:      data.Active,
	}).Error
	if err != nil {
		return &models.UserRole{}, err
	}
	return data, nil
}

func (r *Repository) Delete(id uint64) (int64, error) {
	db := r.DB.Model(&models.UserRole{}).Where("id = ?", id).Take(&models.UserRole{}).Delete(&models.UserRole{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("userRole not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
