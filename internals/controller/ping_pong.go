package controller

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

// PingPongController ...
type PingPongController struct {
	Validator *utils.CustomValidator
}

// StartPing ...
func (ctrl *PingPongController) StartPing(ctx context.Context, req *apiV1.PingPong) (*apiV1.PingPong, error) {
	if err := ctrl.Validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

// NewPingPongController ...
func NewPingPongController(validator *utils.CustomValidator) *PingPongController {
	return &PingPongController{
		Validator: validator,
	}
}
