package models

import "github.com/jinzhu/gorm"

type ParmanentAddress struct {
	Country     string
	State       string
	Street      string
	WardNumber  uint
	HouseNumber string
	PostalCode  string
}

type TemporaryAddress struct {
	TCountry     string
	TState       string
	TStreet      string
	TWardNumber  uint
	THouseNumber string
	TPostalCode  string
}

type Address struct {
	gorm.Model
	ParmanentAddress
	TemporaryAddress
	SID     uint
	Student Student
}
