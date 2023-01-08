package token

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID string, err error)
}
