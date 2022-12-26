package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/users/internals/service/users"
)

type Wrapper struct {
	dig.In  `name:"wrapperUsers"`
	Service service.Service
}
