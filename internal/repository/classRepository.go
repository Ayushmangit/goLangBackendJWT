package repository

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
)

type ClassRepository interface {
	CreateClass(ctx context.Context, args sqlc.CreateClassParams) (sqlc.Class, error)
}

type classRepo struct {
	queries *sqlc.Queries
}

func NewClassRepository(q *sqlc.Queries) ClassRepository {
	return &classRepo{
		queries: q,
	}
}

func (r *classRepo) CreateClass(ctx context.Context, args sqlc.CreateClassParams) (sqlc.Class, error) {
	return r.queries.CreateClass(ctx, args)
}
