package token

import (
	"context"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/token"
)

func (s *LoginService) Update(ctx context.Context, request *model.FitterUpdateToken) (string, error) {
	makeFilter := s.makeFilterTokenUpdate(request)

	if request.LogoutRequest == "logout" {
		tokenUpdate := &entity.Token{
			Status: "Inactive",
		}
		err := s.repository.Update(makeFilter, tokenUpdate)

		return "Logout successful", err
	} else {
		return "Please type 'logout' to logout", nil
	}

}

func (s *LoginService) makeFilterTokenUpdate(filters *model.FitterUpdateToken) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.LogoutRequest) > 0 {
		output["status"] = "Active"
	}

	return output

}
