package users

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var hmacSampleSecret []byte

func (s *RegisterService) Get(ctx context.Context, request *model.FitterReadUsers) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	makeFilter := s.makeFilterUsers(request)
	users := &entity.Users{}
	errx := s.repository.Find(makeFilter, users)
	passwordInput := request.Password

	//compare password login with database
	errCompare := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(passwordInput))
	if errCompare == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_name": request.User_name,
			//"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		return "login successful", errx

	} else {
		return "login failed", nil

	}

	//return &model.ReadResponseUsers{
	//	User_name: users.UserName,
	//	Password:  users.Password,
	//}, err
}

func (s *RegisterService) makeFilterUsers(filters *model.FitterReadUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	if len(filters.Password) > 0 {
		output["password"] = filters.Password
	}

	return output

}
