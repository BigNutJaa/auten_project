package token

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type LoginService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &LoginService{
		repository: r,
	}
}
