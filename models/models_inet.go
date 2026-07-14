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
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	LineID       string `json:"lineid"`
	PhoneID      string `json:"phoneid"`
	BusinessType string `json:"businesstype"`
	WebsiteName  string `json:"websitename"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}
