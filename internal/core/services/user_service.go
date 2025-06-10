package services

import (
	"context"
	"errors"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
	"github.com/AgsOtero/event-ticket-api/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{userRepository: repo}
}

func (s *userService) GetById(ctx context.Context, id int64) (domain.User, error) {
	getUser, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	if getUser.ID == 0 {
		return domain.User{}, errors.New("user not found")
	}
	return getUser, nil
}

func (s *userService) Register(ctx context.Context, email, password, name, surname, phone string) (domain.User, error) {
	existingUser, err := s.userRepository.FindByEmail(ctx, email)
	if err == nil && existingUser.ID != 0 {
		return domain.User{}, errors.New("user with this email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New("could not hash password")
	}
	newUser := domain.User{
		Email:        email,
		Name:         name,
		Surname:      surname,
		PasswordHash: string(hashedPassword),
		Phone:        phone,
	}
	createdUser, err := s.userRepository.Save(ctx, newUser)
	if err != nil {
		return domain.User{}, errors.New("could not save user")
	}
	return createdUser, nil
}
