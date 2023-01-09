package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/users/internals/service/products"
)

type Wrapper struct {
	dig.In  `name:"wrapperProducts"`
	Service service.Service
}
