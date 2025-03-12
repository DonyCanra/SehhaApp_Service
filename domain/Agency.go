package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agency struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Type        string    `gorm:"type:varchar(20);not null;check:type IN ('Klinik', 'Rumah Sakit')" json:"type"`
	Address     string    `gorm:"type:text;not null" json:"address"`
	PhoneNumber string    `gorm:"type:varchar(20);not null; unique" json:"phoneNumber"`
	Email       string    `gorm:"type:varchar(100);unique" json:"email"`
	Website     string    `gorm:"type:varchar(255)" json:"website"`
	Status      bool      `gorm:"type:bool" json:"status"`
	Users       []User    `gorm:"foreignKey:AgencyID" json:"users"`
	BaseAuditEntry
}

// Generate UUID sebelum data disimpan
func (a *Agency) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

// TableName menentukan nama tabel di database
func (a *Agency) TableName() string {
	return "t_agency"
}
