package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/users/internals/service/token"
)

type Wrapper struct {
	dig.In  `name:"wrapperToken"`
	Service service.Service
}
