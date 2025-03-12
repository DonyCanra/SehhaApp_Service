package dto

import "github.com/google/uuid"

type ReqRegisterDTO struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=6"`
	AgencyId uuid.UUID `json:"agencyId"`
	Role     string    `json:"role,omitempty"`
}
