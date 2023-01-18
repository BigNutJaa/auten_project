package wrapper

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/products"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *model.FitterUpdateProducts) (string, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Stock.Update")
	defer sp.Finish()

	id, err := wrp.Service.Update(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
