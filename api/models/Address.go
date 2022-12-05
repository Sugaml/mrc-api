package models

import "github.com/jinzhu/gorm"

type ParmanentAddress struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street      string `json:"street"`
	WardNumber  uint   `json:"ward_number"`
	HouseNumber string `json:"house_number"`
	PostalCode  string `json:"postal_code"`
}

type TemporaryAddress struct {
	TCountry     string `json:"tcountry"`
	TCity        string `json:"tcity"`
	TState       string `json:"tstate"`
	TStreet      string `json:"tstreet"`
	TWardNumber  uint   `json:"tward_number"`
	THouseNumber string `json:"thouse_number"`
	TPostalCode  string `json:"tpostal_code"`
}

type Address struct {
	gorm.Model
	StudentID uint `json:"student_id"`
	ParmanentAddress
	TemporaryAddress
}
