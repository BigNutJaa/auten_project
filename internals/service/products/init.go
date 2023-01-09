package products

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type ProductsService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &ProductsService{
		repository: r,
	}
}
