package repository

import (
	"sugam-project/api/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type IDepartment interface {
	FindAllDepartment(db *gorm.DB) (*[]models.Department, error)
	FindbyId(db *gorm.DB, cid uint) (*models.Department, error)
	SaveDepartment(db *gorm.DB, Department *models.Department) (*models.Department, error)
	UpdateDepartment(db *gorm.DB, department *models.Department, cid uint) (*models.Department, error)
	DeleteDepartment(db *gorm.DB, cid uint) (int64, error)
}

type DepartmentRepo struct {
}

func NewDepartmentRepo() IDepartment {
	return &DepartmentRepo{}
}

func NewDepartment(Department models.Department) *models.Department {
	return &models.Department{
		Name:        Department.Name,
		Description: Department.Description,
		IsActive:    Department.IsActive,
	}
}

func (cr *DepartmentRepo) FindAllDepartment(db *gorm.DB) (*[]models.Department, error) {
	Departments := &[]models.Department{}
	err := db.Model(&models.Department{}).Find(&Departments).Error
	if err != nil {
		log.Error().AnErr("Department find error ::", err)
		return nil, err
	}
	return Departments, nil
}

func (cr *DepartmentRepo) FindbyId(db *gorm.DB, cid uint) (*models.Department, error) {
	Department := &models.Department{}
	err := db.Model(models.Department{}).Where("id = ?", cid).Take(&Department).Error
	if err != nil {
		return &models.Department{}, err
	}
	return Department, nil
}

func (cr *DepartmentRepo) SaveDepartment(db *gorm.DB, Department *models.Department) (*models.Department, error) {
	err := db.Model(&models.Department{}).Create(&Department).Error
	if err != nil {
		log.Error().AnErr("Department save error ::", err)
		return nil, err
	}
	return Department, nil
}

func (cr *DepartmentRepo) UpdateDepartment(db *gorm.DB, department *models.Department, cid uint) (*models.Department, error) {
	data := &models.Department{}
	err := db.Model(&models.Department{}).Where("id = ?", cid).Updates(department).Take(data).Error
	if err != nil {
		return &models.Department{}, err
	}
	return data, nil
}

func (cr *DepartmentRepo) DeleteDepartment(db *gorm.DB, cid uint) (int64, error) {
	result := db.Model(&models.Course{}).Where("id = ?", cid).Delete(&models.Course{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
