package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sehha/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	SaveAndUpdateUser(user domain.User) (domain.User, error)
	GetUserByAgencyId(agencyId uuid.UUID) ([]domain.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func (u userConnection) GetUserByAgencyId(agencyId uuid.UUID) ([]domain.User, error) {
	//TODO implement me
	var data []domain.User

	if err := u.connection.Where("agency_id = ?", agencyId).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (u userConnection) SaveAndUpdateUser(user domain.User) (domain.User, error) {
	if user.AgencyID == uuid.Nil {
		return user, errors.New("agency id is required")
	}

	tx := u.connection.Begin()
	if user.ID == 0 {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return user, err
		}
	} else {
		// Jika ID sudah ada, lakukan update
		if err := tx.Model(&user).Updates(user).Error; err != nil { // Pakai &user
			tx.Rollback()
			return user, err
		}
	}

	tx.Commit()
	return user, nil
}

func (u userConnection) GetUserByEmail(email string) (domain.User, error) {
	//TODO implement me
	var data domain.User

	if err := u.connection.Where("email = ?", email).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil

}

func (u userConnection) GetAll() ([]domain.User, error) {
	//TODO implement me
	var data []domain.User
	if err := u.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}
