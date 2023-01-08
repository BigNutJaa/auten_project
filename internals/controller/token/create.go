package token

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/token"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.TokenCreateRequest) (*apiV1.TokenCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Token.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		User_name: request.GetUserName(),
		Password:  request.GetPassword(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.TokenCreateResponse{Result: string(id)}, nil
}
