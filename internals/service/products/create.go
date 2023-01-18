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
		fmt.Println("User in use :", claims["user_name"])
		input := &entity.Products{
			Name:   request.Name,
			Detail: request.Detail,
			Qty:    request.Qty,
		}
		errx := s.repository.Create(input)

		postSuccess := "Post success, ID:" + strconv.Itoa(input.ID)
		return postSuccess, errx

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

func (s *ProductsService) makeFilterToken(filters string) (output map[string]interface{}) {
	output = make(map[string]interface{})
	if len(filters) > 0 {
		output["token"] = filters
	}
	return output
}

//func (s *ProductsService) makeFilterExpire(filters string) (output map[string]interface{}) {
//	output = make(map[string]interface{})
//	if len(filters) > 0 {
//		output["status"] = filters
//	}
//	return output
//}
