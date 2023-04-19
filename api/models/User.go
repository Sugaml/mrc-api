package models

import (
	"errors"
	"unicode"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Image         string `json:"image"`
	IsAdmin       bool   `json:"is_admin"`
	Active        bool   `json:"active"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
}

type UserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `gorm:"gender" json:"gender"`
	Role      string `json:"role"`
	Image     string `json:"image"`
}

func NewUser(req *UserRequest) *User {
	return &User{
		Email:    req.Email,
		Username: req.Email,
		Password: req.Password,
		Role:     req.Role,
		Image:    req.Image,
	}
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfrimPassword string `json:"confirm_password"`
}

type SetPasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfrimPassword string `json:"confirm_password"`
}

func (user *User) Prepare() {
	user.ID = 0
	user.Active = true
	user.IsAdmin = false
}

func (user *User) Validate() error {
	if user.Email == "" {
		return errors.New("required email")
	}
	if user.Password == "" {
		return errors.New("required password")
	}
	if len(user.Password) < 8 {
		return errors.New("password should be atleast 8 characters")
	}
	return nil
}

func (user *User) LoginValidate() error {
	if user.Username == "" {
		return errors.New("required username")
	}
	if user.Password == "" {
		return errors.New("required password")
	}
	return nil
}

func (req *ChangePasswordRequest) Validate() (string, error) {
	if req.OldPassword == "" {
		return "", ErrRequiredOldPassword
	}
	if req.OldPassword == req.NewPassword {
		return "", ErrSamePasswordAsOld
	}
	return validateChangePassword(req.NewPassword, req.ConfrimPassword)
}

func validateChangePassword(newPassword, confrimPassword string) (string, error) {
	if newPassword == "" {
		return "", ErrRequiredNewPassword
	}
	if confrimPassword == "" {
		return "", ErrRequiredConfrimPassword
	}
	if newPassword != confrimPassword {
		return "", ErrPasswordMismatch
	}
	err := PasswordValidate(newPassword)
	if err != nil {
		return "", err
	}
	return newPassword, nil
}

func PasswordValidate(password string) error {
	// Check password length
	if len(password) < 8 {
		return ErrPasswordMinLength
	}

	// Check for at least one uppercase letter
	hasUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return ErrPasswordUpperCaseLetter
	}

	// Check for at least one lowercase letter
	hasLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
			break
		}
	}
	if !hasLower {
		return ErrPasswordLowerCaseLetter
	}

	// Check for at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return ErrPasswordOneDigit
	}

	// Check for at least one special character
	hasSpecial := false
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return ErrPasswordSpecialCharacter
	}

	return nil
}
