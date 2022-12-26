package users

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID string, err error)
	Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error)
}
