package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IEnrollment interface {
	FindAllEnrollment(db *gorm.DB) (*[]models.Enrollment, error)
	FindbyId(db *gorm.DB, cid uint) (*models.Enrollment, error)
	SaveEnrollment(db *gorm.DB, enrollment *models.Enrollment) (*models.Enrollment, error)
	UpdateEnrollment(db *gorm.DB, enrollment *models.Enrollment, cid uint) (*models.Enrollment, error)
	DeleteEnrollment(db *gorm.DB, cid uint) (int64, error)
}

type EnrollmentRepo struct {
}

func NewEnrollmentRepo() IEnrollment {
	return &EnrollmentRepo{}
}

func NewEnrollment(enrollment models.Enrollment) *models.Enrollment {
	return &models.Enrollment{}
}

func (cr *EnrollmentRepo) FindAllEnrollment(db *gorm.DB) (*[]models.Enrollment, error) {
	Enrollments := &[]models.Enrollment{}
	err := db.Model(&models.Enrollment{}).Find(&Enrollments).Error
	if err != nil {
		log.Error().AnErr("Enrollment find error ::", err)
		return nil, err
	}
	return Enrollments, nil
}

func (cr *EnrollmentRepo) FindbyId(db *gorm.DB, cid uint) (*models.Enrollment, error) {
	Enrollment := &models.Enrollment{}
	err := db.Model(models.Enrollment{}).Where("id = ?", cid).Take(&Enrollment).Error
	if err != nil {
		return &models.Enrollment{}, err
	}
	return Enrollment, nil
}

func (cr *EnrollmentRepo) SaveEnrollment(db *gorm.DB, enrollment *models.Enrollment) (*models.Enrollment, error) {
	err := db.Model(&models.Enrollment{}).Create(&enrollment).Error
	if err != nil {
		log.Error().AnErr("Enrollment save error ::", err)
		return nil, err
	}
	return enrollment, nil
}

func (cr *EnrollmentRepo) UpdateEnrollment(db *gorm.DB, enrollment *models.Enrollment, cid uint) (*models.Enrollment, error) {
	data := &models.Enrollment{}
	err := db.Model(&models.Enrollment{}).Where("id = ?", cid).Updates(enrollment).Take(data).Error
	if err != nil {
		return &models.Enrollment{}, err
	}
	return data, nil
}

func (cr *EnrollmentRepo) DeleteEnrollment(db *gorm.DB, cid uint) (int64, error) {
	result := db.Model(&models.Course{}).Where("id = ?", cid).Delete(&models.Course{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
