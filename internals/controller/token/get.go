package token

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Get(ctx context.Context, request *apiV1.TokenGetRequest) (*apiV1.TokenGetResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.token.Get",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	tokenDatas, err := c.service.Get(ctx, &model.FitterReadToken{
		TokenLogout: request.GetTokenLogout(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.TokenGetResponse{Result: string(tokenDatas)}, nil
}
