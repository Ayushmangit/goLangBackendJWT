package repository

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
)

type SubjectRepository interface {
	CreateSubject(ctx context.Context, args sqlc.CreateSubjectParams) (sqlc.Subject, error)
}

type subjectRepo struct {
	queries *sqlc.Queries
}

func NewSubjectRepository(q *sqlc.Queries) SubjectRepository {
	return &subjectRepo{
		queries: q,
	}
}

func (r *subjectRepo) CreateSubject(ctx context.Context, args sqlc.CreateSubjectParams) (sqlc.Subject, error) {
	return r.queries.CreateSubject(ctx, args)
}
