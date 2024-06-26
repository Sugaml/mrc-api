package repository

import (
	"github.com/Sugaml/mrc-api/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IStudent interface {
	FindAllStudent(db *gorm.DB) (*[]models.Student, error)
	FindbyId(db *gorm.DB, cid uint) (*models.Student, error)
	FindbyUserId(db *gorm.DB, uid uint) (*models.Student, error)
	SaveStudent(db *gorm.DB, Student *models.Student) (*models.Student, error)
	UpdateStudent(db *gorm.DB, Student *models.Student, cid uint) (*models.Student, error)
	UpdateStudentStatus(db *gorm.DB, sid uint, status bool) (*models.Student, error)
	DeleteStudent(db *gorm.DB, cid uint) (int64, error)
}

type StudentRepo struct {
}

func NewStudentRepo() IStudent {
	return &StudentRepo{}
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
	err := db.Model(models.Student{}).Preload("Course").Where("id = ?", cid).Take(&data).Error
	if err != nil {
		return &models.Student{}, err
	}
	return data, nil
}

func (cr *StudentRepo) FindbyUserId(db *gorm.DB, uid uint) (*models.Student, error) {
	data := &models.Student{}
	err := db.Model(models.Student{}).Preload("Course").Preload("User").Where("user_id = ?", uid).Take(&data).Error
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
	req := &models.Student{}
	if data.FirstName != "" {
		req.FirstName = data.FirstName
	}
	if data.LastName != "" {
		req.LastName = data.LastName
	}
	if data.Email != "" {
		req.Email = data.Email
	}
	if data.DOB != "" {
		req.DOB = data.DOB
	}
	if data.MobileNumber != "" {
		req.MobileNumber = data.MobileNumber
	}
	if data.CID != 0 {
		req.CID = data.CID
	}
	if data.UserId != 0 {
		req.UserId = data.UserId
	}
	if data.Gender != "" {
		req.Gender = data.Gender
	}
	if data.ParanetMobile != "" {
		req.ParanetMobile = data.ParanetMobile
	}
	if data.ParanetName != "" {
		req.ParanetName = data.ParanetName
	}
	if data.ParentRelation != "" {
		req.ParentRelation = data.ParentRelation
	}
	err := db.Model(&models.Student{}).Where("id = ?", cid).Updates(req).Take(data).Error
	if err != nil {
		return &models.Student{}, err
	}
	return data, nil
}

func (sr *StudentRepo) UpdateStudentStatus(db *gorm.DB, sid uint, status bool) (*models.Student, error) {
	data := &models.Student{}
	err := db.Model(&models.Student{}).Where("id = ?", sid).UpdateColumn(map[string]interface{}{
		"is_approved": status,
	}).Take(data).Error
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
