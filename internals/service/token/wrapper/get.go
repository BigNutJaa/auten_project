package wrapper

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Get(ctx context.Context, input *model.FitterReadToken) (string, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Token.Get")
	defer sp.Finish()

	id, err := wrp.Service.Get(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
