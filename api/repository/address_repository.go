package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IStudentAddress interface {
	FindStudentAddressbyId(db *gorm.DB, sid uint) (*models.Address, error)
	SaveStudentAddress(db *gorm.DB, address *models.Address) (*models.Address, error)
	UpdateStudentAddress(db *gorm.DB, address *models.Address, cid uint) (*models.Address, error)
	DeleteStudentAddress(db *gorm.DB, sid uint) (int64, error)
}

type StudentAddressRepo struct {
}

func NewStudentAddressRepo() IStudentAddress {
	return &StudentAddressRepo{}
}

func NewStudentAddress(data models.Student) *models.Student {
	return &models.Student{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}
}

func (cr *StudentAddressRepo) FindStudentAddressbyId(db *gorm.DB, sid uint) (*models.Address, error) {
	data := &models.Address{}
	err := db.Model(models.Address{}).Where("student_id = ?", sid).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *StudentAddressRepo) SaveStudentAddress(db *gorm.DB, data *models.Address) (*models.Address, error) {
	err := db.Model(&models.Address{}).Create(&data).Error
	if err != nil {
		log.Error().AnErr("student address data save error ::", err)
		return nil, err
	}
	return data, nil
}

func (cr *StudentAddressRepo) UpdateStudentAddress(db *gorm.DB, address *models.Address, cid uint) (*models.Address, error) {
	err := db.Model(&models.Address{}).Where("id = ?", cid).Updates(address).Take(address).Error
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (cr *StudentAddressRepo) DeleteStudentAddress(db *gorm.DB, sid uint) (int64, error) {
	result := db.Model(&models.Address{}).Where("id = ?", sid).Delete(&models.Address{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
