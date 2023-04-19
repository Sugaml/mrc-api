package models

import "errors"

//Email Require Error
var ErrEmailRequired = errors.New("required email")
var ErrFirstNameRequired = errors.New("required firstname")
var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordRequired = errors.New("required password")
var ErrParentIDRequired = errors.New("required parent id")
var ErrRequiredOldPassword = errors.New("required old passowd")
var ErrSamePasswordAsOld = errors.New("required different than old passowd")
var ErrRequiredNewPassword = errors.New("required new passowd")
var ErrRequiredConfrimPassword = errors.New("required Conformation passowd")
var ErrPasswordMinLength = errors.New("password should be atleast 6 characters")
var ErrPasswordUpperCaseLetter = errors.New("password must contain at least one uppercase letter")
var ErrPasswordLowerCaseLetter = errors.New("password must contain at least one lowercase letter")
var ErrPasswordSpecialCharacter = errors.New("password must contain at least one special character")
var ErrPasswordOneDigit = errors.New("password must contain at least one digit")
var ErrPasswordMismatch = errors.New("password provided do not match")
var ErrPasswordHashGenerate = errors.New("unable to generate hash password")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid model")
