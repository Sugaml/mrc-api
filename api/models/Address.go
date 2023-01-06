package models

import "github.com/jinzhu/gorm"

type ParmanentAddress struct {
	Country      string `json:"country"`
	Provience    string `json:"provience"`
	District     string `json:"district"`
	City         string `json:"city"`
	Municipality string `json:"municipality"`
	Street       string `json:"street"`
	WardNumber   uint   `json:"ward_number"`
	HouseNumber  string `json:"house_number"`
	PostalCode   string `json:"postal_code"`
}

type TemporaryAddress struct {
	TCountry      string `json:"tcountry"`
	TProvience    string `json:"tprovience"`
	TDistrict     string `json:"tdistrict"`
	TCity         string `json:"tcity"`
	TMunicipality string `json:"tmunicipality"`
	TStreet       string `json:"tstreet"`
	TWardNumber   uint   `json:"tward_number"`
	THouseNumber  string `json:"thouse_number"`
	TPostalCode   string `json:"tpostal_code"`
}

type Address struct {
	gorm.Model
	StudentID uint `json:"student_id"`
	ParmanentAddress
	TemporaryAddress
}
