package domain

import "github.com/google/uuid"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AgencyID    uuid.UUID `gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agency_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	Role        string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	EmailVerify bool      `gorm:"type:bool" json:"email_verify"`
	BaseAuditEntry
}

func (User) TableName() string {
	return "t_user"
}
