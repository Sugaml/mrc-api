package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IStudentFile interface {
	FindStudenFilebyId(db *gorm.DB, sid uint) (*models.StudentFile, error)
	SaveStudentFile(db *gorm.DB, sfile *models.StudentFile) (*models.StudentFile, error)
	UpdateStudentFile(db *gorm.DB, sfile *models.StudentFile, cid uint) (*models.StudentFile, error)
	DeleteStudentFile(db *gorm.DB, sid uint) (int64, error)
}

type StudentFileRepo struct {
}

func NewStudentFileRepo() IStudentFile {
	return &StudentFileRepo{}
}

func NewStudentFile(data models.StudentFile) *models.StudentFile {
	return &models.StudentFile{}
}

func (cr *StudentFileRepo) FindStudenFilebyId(db *gorm.DB, sid uint) (*models.StudentFile, error) {
	data := &models.StudentFile{}
	err := db.Model(models.StudentFile{}).Preload("Student").Where("student_id = ?", sid).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *StudentFileRepo) SaveStudentFile(db *gorm.DB, sfile *models.StudentFile) (*models.StudentFile, error) {
	err := db.Model(&models.StudentFile{}).Create(sfile).Error
	if err != nil {
		log.Error().AnErr("student file data save error ::", err)
		return nil, err
	}
	return sfile, nil
}

func (cr *StudentFileRepo) UpdateStudentFile(db *gorm.DB, sfile *models.StudentFile, cid uint) (*models.StudentFile, error) {
	err := db.Model(&models.StudentFile{}).Where("id = ?", cid).Updates(sfile).Take(sfile).Error
	if err != nil {
		return nil, err
	}
	return sfile, nil
}

func (cr *StudentFileRepo) DeleteStudentFile(db *gorm.DB, sid uint) (int64, error) {
	result := db.Model(&models.StudentFile{}).Where("id = ?", sid).Delete(&models.StudentFile{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
