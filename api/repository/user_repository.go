package repository

import (
	"html"
	"regexp"
	"strings"
	"sugam-project/api/models"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	FindAll(db *gorm.DB) (*[]models.User, error)
	FindbyId(db *gorm.DB, uid uint) (*models.User, error)
	FindbyUsername(db *gorm.DB, username string) (*models.User, error)
	Save(db *gorm.DB, user *models.User) (*models.User, error)
	Update(db *gorm.DB, user *models.User, uid uint) (*models.User, error)
	Delete(db *gorm.DB, uid uint) (int64, error)
}

type UserRepo struct{}

func NewUserRepo() IUser {
	return &UserRepo{}
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func isHashed(password string) bool {
	return len(password) == 60
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// func (u *User) BeforeSave() error {
// 	if !isHashed(u.Password) {
// 		hashedPassword, err := Hash(u.Password)
// 		if err != nil {
// 			return err
// 		}
// 		u.Password = string(hashedPassword)
// 	}
// 	return nil
// }

func Prepares(u *models.User) {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	// u.Company = html.EscapeString(strings.TrimSpace(u.Company))
	// u.Designation = html.EscapeString(strings.TrimSpace(u.Designation))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.IsAdmin = false
}
func ValidateName(name string) bool {
	valid, _ := regexp.Match("^\\w+([\\s-_]\\w+)*$", []byte(name))
	return valid
}

// func (u *models.User) Validate(action string) error {
// 	switch strings.ToLower(action) {
// 	case "update":
// 		if u.FirstName == "" {
// 			return errors.New("Required FirstName")
// 		}
// 		if u.LastName == "" {
// 			return errors.New("Required LastName")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		if !ValidateName(u.FirstName) {
// 			return errors.New("Allowed alphanumeric, underscore, hyphen and space only.")
// 		}
// 		return nil
// 	case "login":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	case "forgotpassword":
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if u.Email != "" {
// 			if checkmail.ValidateFormat(u.Email) != nil {
// 				return errors.New("Invalid Email")
// 			}
// 		}
// 		return nil
// 	default:
// 		if u.FirstName == "" {
// 			return errors.New("Required FirstName")
// 		}
// 		if u.LastName == "" {
// 			return errors.New("Required LastName")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	}
// }

// func NewUser(user models.User) *models.User {
// 	return &models.User{
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		Username:  user.Username,
// 		Active:    user.Active,
// 		IsAdmin:   user.IsAdmin,
// 	}
// }

func (cr *UserRepo) FindAll(db *gorm.DB) (*[]models.User, error) {
	user := &[]models.User{}
	err := db.Model(&models.User{}).Find(&user).Error
	if err != nil {
		log.Error().AnErr("course save error ::", err)
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) FindbyId(db *gorm.DB, uid uint) (*models.User, error) {
	user := &models.User{}
	err := db.Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) FindbyUsername(db *gorm.DB, username string) (*models.User, error) {
	user := &models.User{}
	err := db.Model(models.User{}).Where("email = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) Save(db *gorm.DB, user *models.User) (*models.User, error) {
	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Error().AnErr("course save error ::", err)
		return nil, err
	}
	return user, nil
}

func (cr *UserRepo) Update(db *gorm.DB, user *models.User, uid uint) (*models.User, error) {
	data := &models.User{}
	err := db.Model(&models.User{}).Where("id = ?", uid).Updates(user).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cr *UserRepo) Delete(db *gorm.DB, uid uint) (int64, error) {
	result := db.Model(&models.User{}).Where("id = ?", uid).Delete(&models.User{})
	if result.Error != nil {
		return 0, db.Error
	}
	return result.RowsAffected, nil
}
