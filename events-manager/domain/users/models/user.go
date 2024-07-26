package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email string `json:"email,omitempty" binding:"required"`

	Name string `json:"name,omitempty" binding:"required"`

	Password string `json:"password,omitempty" binding:"required"`

	PurposeOfUse string `json:"purposeOfUse,omitempty" binding:"required"`
}
