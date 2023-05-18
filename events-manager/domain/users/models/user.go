package models

type User struct {
	Id string `json:"__id,omitempty"`

	Email string `json:"summary,omitempty"`

	Password string `json:"description,omitempty"`
}
