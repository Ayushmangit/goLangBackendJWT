package service

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/repository"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(
		ctx context.Context,
		userID uuid.UUID,
		fullName string,
		rollNumber string,
		classID uuid.UUID,
	) (sqlc.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(r repository.StudentRepository) StudentService {
	return &studentService{
		repo: r,
	}
}

func (s *studentService) Create(
	ctx context.Context,
	userID uuid.UUID,
	fullName string,
	rollNumber string,
	classID uuid.UUID,
) (sqlc.Student, error) {

	payload := sqlc.CreateStudentParams{
		ID:         uuid.New(),
		UserID:     userID,
		FullName:   fullName,
		RollNumber: rollNumber,
		ClassID:    classID,
	}

	return s.repo.CreateStudent(ctx, payload)
}
