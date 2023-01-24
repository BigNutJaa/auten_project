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
	"time"
)

var hmacSampleSecret []byte

func (s *LoginService) Create(ctx context.Context, request *model.Request) (string, error) {

	// Find user&password from database
	passwordLogin := request.Password
	userLogin := request.User_name
	checkMatching := s.makeFilterPasswordExist(request)
	users := &entity.Users{}
	err := s.repository.Find(checkMatching, users)
	resultCheck, _ := &model2.ReadResponseUsers{
		User_name: users.UserName,
		Password:  users.Password,
		Role_code: users.RoleCode,
	}, err
	databasePassword := resultCheck.Password
	databaseUser := resultCheck.User_name

	fmt.Println("check user login:", userLogin)

	// Compare login - database
	errCompare := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(passwordLogin))
	if userLogin != databaseUser {
		return "Login failed! :Username does not exist", nil

	} else if errCompare == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		//hmacSampleSecret = []byte("my_secret_key")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_name": request.User_name,
			"role_code": resultCheck.Role_code,
			"exp":       time.Now().Add(time.Minute * 10).Unix(),
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
		loginSuccess := "Login successful! Token is: " + input.Token
		return loginSuccess, errx

	} else {
		return "Login failed! : Password incorrect", nil

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
