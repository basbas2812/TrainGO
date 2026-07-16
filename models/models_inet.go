package models

import "gorm.io/gorm"

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Register struct {
	Email        string `json:"email"        validate:"required,email"`
	Username     string `json:"username"     validate:"required,username"`
	Password     string `json:"password"     validate:"required,min=6,max=20"`
	LineID       string `json:"lineid"       validate:"required,lineid"`
	PhoneID      string `json:"phoneid"      validate:"required,min=9,max=10,numeric"`
	BusinessType string `json:"businesstype" validate:"required,business_allowed"`
	WebsiteName  string `json:"websitename"  validate:"required,min=2,max=30,websitename"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name"`
	Address  string `json:"address"`
	Tel      string `json:"telephone"`
	EmpCount int    `json:"emp_count"`
	IsActive bool   `json:"is_active"`
}

type DogColor struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Color string `json:"color"`
}
