package service

import (
	"context"
	"errors"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, email, password string) (sqlc.User, error)
	Login(ctx context.Context, email, password string) (sqlc.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (sqlc.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) Register(ctx context.Context, email, password string) (sqlc.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, err
	}

	payload := sqlc.CreateUserParams{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	user, err := s.repo.CreateUser(ctx, payload)
	if err != nil {
		return sqlc.User{}, errors.New("user already exists")
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (sqlc.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return sqlc.User{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return sqlc.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (sqlc.User, error) {
	return s.repo.FindByID(ctx, id)
}
