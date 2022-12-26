package users

import (
	"github.com/BigNutJaa/users/internals/service/users/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
