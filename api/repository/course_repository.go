package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type ICourse interface {
	FindAllCourse(db *gorm.DB) (*[]models.Course, error)
	FindbyId(db *gorm.DB, cid uint) (*models.Course, error)
	SaveCourse(db *gorm.DB, course *models.Course) (*models.Course, error)
	UpdateCourse(db *gorm.DB, course *models.Course, cid uint) (*models.Course, error)
	DeleteCourse(db *gorm.DB, cid uint) (int64, error)
}

type CourseRepo struct {
}

func NewCourseRepo() ICourse {
	return &CourseRepo{}
}

func NewCourse(course models.Course) *models.Course {
	return &models.Course{
		Name:        course.Name,
		Discription: course.Discription,
		Fee:         course.Fee,
		Duration:    course.Duration,
		//Affiliatedby: course.Affiliatedby,
		IsActive: course.IsActive,
	}
}

func (cr *CourseRepo) FindAllCourse(db *gorm.DB) (*[]models.Course, error) {
	courses := &[]models.Course{}
	err := db.Model(&models.Course{}).Find(&courses).Error
	if err != nil {
		log.Error().AnErr("course find error ::", err)
		return nil, err
	}
	return courses, nil
}

func (cr *CourseRepo) FindbyId(db *gorm.DB, cid uint) (*models.Course, error) {
	course := &models.Course{}
	err := db.Model(models.Course{}).Where("id = ?", cid).Take(&course).Error
	if err != nil {
		return &models.Course{}, err
	}
	return course, nil
}

func (cr *CourseRepo) SaveCourse(db *gorm.DB, course *models.Course) (*models.Course, error) {
	err := db.Model(&models.Course{}).Create(&course).Error
	if err != nil {
		log.Error().AnErr("course save error ::", err)
		return nil, err
	}
	return course, nil
}

func (cr *CourseRepo) UpdateCourse(db *gorm.DB, course *models.Course, cid uint) (*models.Course, error) {
	data := &models.Course{}
	err := db.Model(&models.Course{}).Where("id = ?", cid).Updates(course).Take(data).Error
	if err != nil {
		return &models.Course{}, err
	}
	return data, nil
}

func (cr *CourseRepo) DeleteCourse(db *gorm.DB, cid uint) (int64, error) {
	result := db.Model(&models.Course{}).Where("id = ?", cid).Delete(&models.Course{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
