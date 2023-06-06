package dtos

type LoginDTO struct {
	Email    string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
