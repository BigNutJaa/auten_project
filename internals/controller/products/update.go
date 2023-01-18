package products

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/products"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/metadata"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.ProductsUpdateRequest) (*apiV1.ProductsUpdateResponse, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	tokenx := md.Get("token")

	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.products.Update",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	productsData, err := c.service.Update(ctx, &model.FitterUpdateProducts{
		Name:      request.GetName(),
		Detail:    request.GetDetail(),
		Id:        request.GetId(),
		QtyUpdate: request.GetQtyUpdate(),
		Token:     tokenx[0],
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.ProductsUpdateResponse{Result: string(productsData)}, nil
}
