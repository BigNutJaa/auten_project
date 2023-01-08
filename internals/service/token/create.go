package token

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/token"
	model2 "github.com/BigNutJaa/users/internals/model/users"
	"golang.org/x/crypto/bcrypt"
)

func (s *LoginService) Create(ctx context.Context, request *model.Request) (string, error) {

	// encrypt password
	//encryptPass := StartEncrypt(request.Password)
	//encryptPass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	// check password matching
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

	fmt.Println("login pass:", passwordLogin)
	fmt.Println("database pass:", databasePassword)

	//if passwordLogin == resultCheckPassword.Password {
	errCompare := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(passwordLogin))
	if userLogin != databaseUser {
		return "Login failed! :Username does not exist", nil

	} else if errCompare == nil {
		//hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//	"user_name": request.User_name,
		//	//"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		//})
		//// Sign and get the complete encoded token as a string using the secret
		//tokenString, err := token.SignedString(hmacSampleSecret)
		//fmt.Println(tokenString, err)

		return "Login success!", nil

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
