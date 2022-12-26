package users

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Get(ctx context.Context, request *apiV1.UsersGetRequest) (*apiV1.UsersGetResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Users.Get",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	usersData, err := c.service.Get(ctx, &model.FitterReadUsers{
		User_name: request.GetUserName(),
		Password:  request.GetPassword(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersGetResponse{
		UserName: usersData.User_name,
		Password: usersData.Password,
	}, nil
}
