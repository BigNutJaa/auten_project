package products

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/products"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID string, err error)
	Update(ctx context.Context, request *model.FitterUpdateProducts) (ID string, err error)
	//Get(ctx context.Context, request *model.FitterReadUsers) (string, error)
}
