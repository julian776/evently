package models

type User struct {
	Email string `json:"email,omitempty" binding:"required"`

	Name string `json:"name,omitempty" binding:"required"`

	Password string `json:"password,omitempty" binding:"required"`

	PurposeOfUse string `json:"purposeOfUse,omitempty" binding:"required"`
}
