package agency

import "github.com/google/uuid"

type RequestAgencyDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Address     string    `json:"address" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Type        string    `json:"type"`
	Website     string    `json:"website"`
	Status      bool      `json:"status"`
}
