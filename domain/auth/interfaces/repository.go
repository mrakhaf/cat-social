package interfaces

import "context"

type Repository interface {
	Login(ctx context.Context, email, password string) (string, error)
}
