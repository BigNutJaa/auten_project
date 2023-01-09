package products

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/products"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.ProductsCreateRequest) (*apiV1.ProductsCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Products.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
		Qty:    request.GetQty(),
	})
	//First_name: request.GetFirstName(),

	if err != nil {
		return nil, err
	}
	return &apiV1.ProductsCreateResponse{Result: string(id)}, nil
}
