package types

import "github.com/google/uuid"

type AuthUser struct {
	ID    uuid.UUID
	Email string
	Role  string
}
