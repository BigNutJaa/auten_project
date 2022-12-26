package users

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type AutenService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &AutenService{
		repository: r,
	}
}
