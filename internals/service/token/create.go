package token

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/token"
	model2 "github.com/BigNutJaa/users/internals/model/users"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var hmacSampleSecret []byte

func (s *LoginService) Create(ctx context.Context, request *model.Request) (string, error) {

	// where user&password
	passwordLogin := request.Password
	userLogin := request.User_name
	checkMatching := s.makeFilterPasswordExist(request)
	users := &entity.Users{}
	err := s.repository.Find(checkMatching, users)
	resultCheck, _ := &model2.ReadResponseUsers{
		User_name: users.UserName,
		Password:  users.Password,
	}, err
	databasePassword := resultCheck.Password
	databaseUser := resultCheck.User_name

	// Compare login - database
	errCompare := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(passwordLogin))
	if userLogin != databaseUser {
		return "Login failed! :Username does not exist", nil

	} else if errCompare == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		//hmacSampleSecret = []byte("my_secret_key")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_name": request.User_name,
			//"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		input := &entity.Token{
			UserName: userLogin,
			Token:    tokenString,
			Status:   "Active",
		}
		errx := s.repository.Create(input)

		return "Login success!", errx

	} else {
		return "Login failed! : Password incorrect", nil
		//input := &entity.Users{
		//	UserName:  request.User_name,
		//	Password:  string(encryptPass),
		//	FirstName: request.First_name,
		//	LastName:  request.Last_name,
		//	Email:     request.Email,
		//	RoleCode:  "U02_R00",
		//}

		//err := s.repository.Create(input)

		//sp.LogKV("Repository result  :", err)

		//return "User register successfully", err
	}
}

func (s *LoginService) makeFilterPasswordExist(filters *model.Request) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	// ** use username to filter only (password in database already encrypt cannot filter)
	//if len(filters.Password) > 0 {
	//	output["password"] = filters.Password
	//}

	return output

}
