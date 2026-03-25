package repository

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
	FindByID(ctx context.Context, id pgtype.UUID) (sqlc.User, error)
}

type userRepo struct {
	queries *sqlc.Queries
}

// dependency injection
func NewUserRepository(q *sqlc.Queries) UserRepository {
	return &userRepo{
		queries: q,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}

func (r *userRepo) FindByID(ctx context.Context, id pgtype.UUID) (sqlc.User, error) {

	user, err := r.queries.FindByID(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
