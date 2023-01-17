package token

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.TokenUpdateRequest) (*apiV1.TokenUpdateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.token.Update",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	tokenData, err := c.service.Update(ctx, &model.FitterUpdateToken{
		LogoutRequest: request.LogoutRequest,
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.TokenUpdateResponse{Result: string(tokenData)}, nil

}
