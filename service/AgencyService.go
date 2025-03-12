package service

import (
	"net/http"
	"sehha/domain"
	"sehha/dto/agency"
	"sehha/logger"
	"sehha/repository"
)

type AgencyService interface {
	SaveAndUpdate(request agency.RequestAgencyDTO) (domain.Agency, error, int)
}

type agencyService struct {
	agencyRepository repository.AgencyRepository
	userRepository   repository.UserRepository
}

func (a agencyService) SaveAndUpdate(request agency.RequestAgencyDTO) (domain.Agency, error, int) {
	// Inisialisasi objek agency
	agencyData := domain.Agency{
		ID:          request.ID,
		Name:        request.Name,
		Email:       request.Email,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Status:      request.Status,
		Type:        request.Type,
	}

	// Simpan atau update agency
	saveAgency, err := a.agencyRepository.SaveAndUpdate(agencyData)
	if err != nil {
		return saveAgency, err, http.StatusInternalServerError
	}

	logger.Info("SaveAndUpdate: Agency Created", saveAgency)

	// Jika agency berhasil dibuat, buat user dengan role ADMIN
	userRequest := domain.User{
		Name:     agencyData.Name,
		Email:    agencyData.Email,
		Role:     "ADMIN",
		Password: "Rahasia123@",
		AgencyID: saveAgency.ID,
	}

	saveDataUser, errSaveUser := a.userRepository.SaveAndUpdateUser(userRequest)
	if errSaveUser != nil {
		return saveAgency, errSaveUser, http.StatusInternalServerError
	}

	// get user by agency id
	getUser, err := a.userRepository.GetUserByAgencyId(saveAgency.ID)
	if err != nil {
		return saveAgency, err, http.StatusInternalServerError
	}

	saveAgency.Users = getUser

	logger.Info("SaveAndUpdate: User Created", saveDataUser)

	return saveAgency, nil, http.StatusOK
}

func NewAgencyService(agencyRepository repository.AgencyRepository, userRepository repository.UserRepository) AgencyService {
	return &agencyService{
		agencyRepository: agencyRepository,
		userRepository:   userRepository,
	}
}
