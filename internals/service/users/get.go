package users

import (
	"context"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
)

func (s *AutenService) Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error) {
	makeFilter := s.makeFilterUsers(request)
	users := &entity.Users{}
	err := s.repository.Find(makeFilter, users)

	return &model.ReadResponseUsers{
		User_name: users.UserName,
		Password:  users.Password,
	}, err
}

func (s *AutenService) makeFilterUsers(filters *model.FitterReadUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	if len(filters.Password) > 0 {
		output["password"] = filters.Password
	}

	return output

}
