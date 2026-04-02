package service

import (
	"context"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
)

type StudentService interface {
	Create(ctx context.Context, rollNo,) (sqlc.Student, error)
}
