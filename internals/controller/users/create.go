package users

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.UsersCreateRequest) (*apiV1.UsersCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Users.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		User_name:  request.GetUserName(),
		Password:   request.GetPassword(),
		First_name: request.GetFirstName(),
		Last_name:  request.GetLastName(),
		Email:      request.GetEmail(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersCreateResponse{Result: string(id)}, nil
}
