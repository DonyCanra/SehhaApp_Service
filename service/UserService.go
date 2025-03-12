package service

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"sehha/domain"
	"sehha/dto"
	"sehha/logger"
	"sehha/repository"
)

type UserService interface {
	ValidateUser(request dto.ReqLoginDTO) (domain.User, error, int)
	SaveUser(request dto.ReqRegisterDTO) (domain.User, error, int)
}

type userService struct {
	userRepository repository.UserRepository
}

func (a *userService) SaveUser(request dto.ReqRegisterDTO) (domain.User, error, int) {
	var userData domain.User

	// Validasi Agency ID
	if request.AgencyId == uuid.Nil {
		return userData, errors.New("invalid agency ID"), http.StatusBadRequest
	}

	// Set data user
	userData = domain.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		Role:        request.Role,
		AgencyID:    request.AgencyId,
		EmailVerify: false,
	}

	// Simpan user
	userData, err := a.userRepository.SaveAndUpdateUser(userData)
	if err != nil {
		return userData, err, http.StatusInternalServerError
	}

	return userData, nil, http.StatusOK
}

func (a *userService) ValidateUser(request dto.ReqLoginDTO) (domain.User, error, int) {
	userData, err := a.userRepository.GetUserByEmail(request.Email)
	logger.Info("ValidateUser", userData)

	// Jika user tidak ditemukan
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userData, errors.New("user not found"), http.StatusNotFound
	}
	if err != nil {
		return userData, err, http.StatusInternalServerError
	}

	// Validasi password
	if userData.Password != request.Password {
		return userData, errors.New("user name or password not match"), http.StatusBadRequest
	}

	return userData, nil, http.StatusOK
}

func NewAuthService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
