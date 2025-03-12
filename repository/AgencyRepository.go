package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sehha/domain"
)

type AgencyRepository interface {
	GetAll() ([]domain.Agency, error)
	GetById(id uuid.UUID) (domain.Agency, error)
	SaveAndUpdate(request domain.Agency) (domain.Agency, error)
	Delete(id uuid.UUID) (domain.Agency, error)
}

type agencyConnection struct {
	connection *gorm.DB
}

func (a agencyConnection) GetAll() ([]domain.Agency, error) {
	//TODO implement me
	var data []domain.Agency
	if err := a.connection.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (a agencyConnection) GetById(id uuid.UUID) (domain.Agency, error) {
	//TODO implement me
	var data domain.Agency
	if id == uuid.Nil {
		return data, errors.New("invalid agency ID")
	}

	if err := a.connection.Where("id = ?", id).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (a agencyConnection) SaveAndUpdate(request domain.Agency) (domain.Agency, error) {
	//TODO implement me
	var existingAgency domain.Agency
	err := a.connection.Where("id = ?", request.ID).First(&existingAgency).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := a.connection.Create(&request).Error; err != nil {
				return existingAgency, err
			}
			return request, nil
		}
		return existingAgency, err
	}

	// Data ditemukan, lakukan update
	if err := a.connection.Model(&existingAgency).Updates(request).Error; err != nil {
		return existingAgency, err
	}

	return existingAgency, nil
}

// Delete agency
func (a agencyConnection) Delete(id uuid.UUID) (domain.Agency, error) {
	var agency domain.Agency
	err := a.connection.Where("id = ?", id).First(&agency).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return agency, errors.New("agency not found")
		}
		return agency, err
	}

	if err := a.connection.Delete(&agency).Error; err != nil {
		return agency, err
	}

	return agency, nil
}

func NewAgencyRepository(db *gorm.DB) AgencyRepository {
	return &agencyConnection{
		connection: db,
	}

}
