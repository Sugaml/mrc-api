package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IStudent interface {
	FindAllStudent(db *gorm.DB) (*[]models.Student, error)
	FindbyId(db *gorm.DB, cid uint) (*models.Student, error)
	SaveStudent(db *gorm.DB, Student *models.Student) (*models.Student, error)
	UpdateStudent(db *gorm.DB, Student *models.Student, cid uint) (*models.Student, error)
	DeleteStudent(db *gorm.DB, cid uint) (int64, error)
}

type StudentRepo struct {
}

func NewStudentRepo() IStudent {
	return &StudentRepo{}
}

func NewStudent(data models.Student) *models.Student {
	return &models.Student{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}
}

func (cr *StudentRepo) FindAllStudent(db *gorm.DB) (*[]models.Student, error) {
	datas := &[]models.Student{}
	err := db.Model(&models.Student{}).Find(&datas).Error
	if err != nil {
		log.Error().AnErr("get students error ::", err)
		return nil, err
	}
	return datas, nil
}

func (cr *StudentRepo) FindbyId(db *gorm.DB, cid uint) (*models.Student, error) {
	data := &models.Student{}
	err := db.Model(models.Student{}).Where("id = ?", cid).Take(&data).Error
	if err != nil {
		return &models.Student{}, err
	}
	return data, nil
}

func (cr *StudentRepo) SaveStudent(db *gorm.DB, data *models.Student) (*models.Student, error) {
	err := db.Model(&models.Student{}).Create(&data).Error
	if err != nil {
		log.Error().AnErr("student data save error ::", err)
		return nil, err
	}
	return data, nil
}

func (cr *StudentRepo) UpdateStudent(db *gorm.DB, data *models.Student, cid uint) (*models.Student, error) {
	err := db.Model(&models.Student{}).Where("id = ?", cid).Updates(data).Take(data).Error
	if err != nil {
		return &models.Student{}, err
	}
	return data, nil
}

func (cr *StudentRepo) DeleteStudent(db *gorm.DB, cid uint) (int64, error) {
	result := db.Model(&models.Student{}).Where("id = ?", cid).Delete(&models.Course{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
