package models

type User struct {
	Email string `json:"email,omitempty" binding:"required"`

	Name string `json:"name,omitempty" binding:"required"`

	Password string `json:"description,omitempty" binding:"required"`

	PurpouseOfUse string `json:"purpouseOfUse,omitempty" binding:"required"`
}
