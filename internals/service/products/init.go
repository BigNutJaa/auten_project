package products

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type ProductsService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &ProductsService{
		repository: r,
	}
}

func (s *ProductsService) makeFilterToken(filters string) (output map[string]interface{}) {
	output = make(map[string]interface{})
	if len(filters) > 0 {
		output["token"] = filters
	}
	return output
}

func Int32toString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
