package repository

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, arg sqlc.CreateStudentParams) (sqlc.Student, error)
}

type studentRepo struct {
	queries *sqlc.Queries
}

func NewStudentRepository(q *sqlc.Queries) StudentRepository {
	return &studentRepo{
		queries: q,
	}
}

func (r *studentRepo) CreateStudent(ctx context.Context, arg sqlc.CreateStudentParams) (sqlc.Student, error) {
	return r.queries.CreateStudent(ctx, arg)
}
