package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IStudentEducation interface {
	FindStudenEducationbyId(db *gorm.DB, sid uint) (*models.StudentEducation, error)
	FindStudenEducationDetail(db *gorm.DB, sid uint) (*[]models.StudentEducation, error)
	SaveStudentEducation(db *gorm.DB, sedu *models.StudentEducation) (*models.StudentEducation, error)
	UpdateStudentEducation(db *gorm.DB, sedu *models.StudentEducation, cid uint) (*models.StudentEducation, error)
	DeleteStudentEducation(db *gorm.DB, sid uint) (int64, error)
}

type StudentEducationRepo struct {
}

func NewStudentEducationRepo() IStudentEducation {
	return &StudentEducationRepo{}
}

func NewStudentEducation(data models.StudentFile) *models.StudentFile {
	return &models.StudentFile{}
}

func (cr *StudentEducationRepo) FindStudenEducationbyId(db *gorm.DB, sid uint) (*models.StudentEducation, error) {
	data := &models.StudentEducation{}
	err := db.Model(models.StudentEducation{}).Where("student_id = ?", sid).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *StudentEducationRepo) FindStudenEducationDetail(db *gorm.DB, sid uint) (*[]models.StudentEducation, error) {
	data := &[]models.StudentEducation{}
	err := db.Model(models.StudentEducation{}).Where("student_id = ?", sid).Find(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *StudentEducationRepo) SaveStudentEducation(db *gorm.DB, sedu *models.StudentEducation) (*models.StudentEducation, error) {
	err := db.Model(&models.StudentEducation{}).Create(sedu).Error
	if err != nil {
		log.Error().AnErr("student file data save error ::", err)
		return nil, err
	}
	return sedu, nil
}

func (cr *StudentEducationRepo) UpdateStudentEducation(db *gorm.DB, sedu *models.StudentEducation, cid uint) (*models.StudentEducation, error) {
	err := db.Model(&models.StudentEducation{}).Where("id = ?", cid).Updates(sedu).Take(sedu).Error
	if err != nil {
		return nil, err
	}
	return sedu, nil
}

func (cr *StudentEducationRepo) DeleteStudentEducation(db *gorm.DB, sid uint) (int64, error) {
	result := db.Model(&models.StudentEducation{}).Where("id = ?", sid).Delete(&models.StudentEducation{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
