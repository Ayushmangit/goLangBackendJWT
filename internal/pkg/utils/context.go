package utils

import (
	"context"
	"errors"

	"github.com/Ayushmangit/goLangBackendJWT/internal/middleware"
	"github.com/Ayushmangit/goLangBackendJWT/internal/types"
)

var ErrUnauthorized = errors.New("unauthorized")

func GetUserFromContext(ctx context.Context) (*types.AuthUser, error) {
	user, ok := ctx.Value(middleware.UserKey).(*types.AuthUser)
	if !ok || user == nil {
		return nil, ErrUnauthorized
	}
	return user, nil
}

func MustGetUser(ctx context.Context) *types.AuthUser {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		panic("user not found in context")
	}
	return user
}
