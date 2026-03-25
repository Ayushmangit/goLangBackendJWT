package utils

import (
	"context"
	"errors"

	"github.com/Ayushmangit/goLangBackendJWT/internal/middleware"
	"github.com/Ayushmangit/goLangBackendJWT/internal/types"
)

func GetUserFromContext(ctx context.Context) (*types.AuthUser, error) {
	user, ok := ctx.Value(middleware.UserKey).(*types.AuthUser)
	if !ok || user == nil {
		return nil, errors.New("unauthorized")
	}
	return user, nil
}
