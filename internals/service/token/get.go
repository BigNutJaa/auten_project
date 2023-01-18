package token

import (
	"context"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/token"
)

func (s *LoginService) Get(ctx context.Context, request *model.FitterReadToken) (string, error) {

	makeFilter := s.makeFilterToken(request)
	token := &entity.Token{}
	err := s.repository.Find(makeFilter, token)
	resultCheck, _ := &model.ReadResponseToken{
		User_name: token.UserName,
		Token:     token.Token,
	}, err
	databaseToken := resultCheck.Token

	if request.TokenLogout == databaseToken {
		tokenUpdate := &entity.Token{
			Status: "Inactive",
		}
		err := s.repository.Update(makeFilter, tokenUpdate)

		return "Logout successful", err
	} else {
		return "Logout failed! incorrect Token!", nil
	}

}

func (s *LoginService) makeFilterToken(filters *model.FitterReadToken) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.TokenLogout) > 0 {
		output["token"] = filters.TokenLogout
	}

	return output

}
