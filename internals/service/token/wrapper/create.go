package wrapper

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/token"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *model.Request) (string, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Token.Create")
	defer sp.Finish()

	id, err := wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
