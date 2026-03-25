package service

import (
	"context"
	"errors"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/utils"
	"github.com/Ayushmangit/goLangBackendJWT/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, email, password string) (sqlc.User, error)
	Login(ctx context.Context, email, password string) (sqlc.User, error)
	GetByID(ctx context.Context, id string) (sqlc.User, error)
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

	var payload sqlc.CreateUserParams
	payload.ID = utils.NewPGUUID() // as PGX and SQLC will not take UUID
	payload.Email = email
	payload.Password = string(hashedPassword)
	user, err := s.repo.CreateUser(ctx, payload)
	if err != nil {
		return sqlc.User{}, errors.New("user already exists")
	}
	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (sqlc.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return sqlc.User{}, errors.New("invalid Credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return sqlc.User{}, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (sqlc.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, err
	}

	pgID := pgtype.UUID{
		Bytes: uid,
		Valid: true,
	}

	return s.repo.FindByID(ctx, pgID)
}
