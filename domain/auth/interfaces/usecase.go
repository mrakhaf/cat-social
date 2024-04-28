package interfaces

import "context"

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
}
