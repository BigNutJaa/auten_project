package wrapper

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/users"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Get(ctx context.Context, input *model.FitterReadUsers) (*model.ReadResponseUsers, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Users.Get")
	defer sp.Finish()

	id, err := wrp.Service.Get(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
