package products

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/products"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
)

func (s *ProductsService) Create(ctx context.Context, request *model.Request) (string, error) {

	//check token before POST
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString := request.Token

	fmt.Println("test login token:", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("User in use :", claims["user_name"], claims["role_code"])
		user_role := claims["role_code"]

		if CheckRoleCreate(user_role) == true {
			input := &entity.Products{
				Name:   request.Name,
				Detail: request.Detail,
				Qty:    request.Qty,
			}
			errx := s.repository.Create(input)

			postSuccess := "Post success, ID:" + strconv.Itoa(input.ID)
			return postSuccess, errx
		} else {
			return "You have no permission to process", nil
		}

	} else {
		makeFilterEXP := s.makeFilterToken(tokenString)
		tokenUpdate := &entity.Token{
			Status: "Expired",
		}
		errz := s.repository.Update(makeFilterEXP, tokenUpdate)
		fmt.Println("Re-check err wording", err)
		return "Token is expired", errz
	}
}
