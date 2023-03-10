package users

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type RegisterService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &RegisterService{
		repository: r,
	}
}
