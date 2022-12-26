package users

import (
	"context"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
	"golang.org/x/crypto/bcrypt"
)

func (s *AutenService) Create(ctx context.Context, request *model.Request) (string, error) {

	// encrypt password
	//encryptPass := StartEncrypt(request.Password)
	encryptPass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	// check user exists
	userExist := request.User_name
	checkExist := s.makeFilterUserExist(userExist)
	users := &entity.Users{}
	err := s.repository.Find(checkExist, users)
	resultCheckExist, _ := &model.ReadResponseUsers{
		User_name: users.UserName,
	}, err

	if userExist == resultCheckExist.User_name {
		return "Register failed: Username already exist", nil
	} else {

		input := &entity.Users{
			UserName:  request.User_name,
			Password:  string(encryptPass),
			FirstName: request.First_name,
			LastName:  request.Last_name,
			Email:     request.Email,
			RoleCode:  "U02_R00",
		}

		err := s.repository.Create(input)

		//sp.LogKV("Repository result  :", err)

		return "User register successfully", err
	}
}

func (s *AutenService) makeFilterUserExist(filters string) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters) > 0 {
		output["user_name"] = filters
	}

	return output

}
