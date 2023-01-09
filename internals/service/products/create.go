package products

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/products"
	model2 "github.com/BigNutJaa/users/internals/model/token"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func (s *ProductsService) Create(ctx context.Context, request *model.Request) (string, error) {
	checkMatching := s.makeFilterToken("Active")
	tokenX := &entity.Token{}
	erry := s.repository.Find(checkMatching, tokenX)
	resultCheckToken, _ := &model2.ReadResponseToken{
		User_name: tokenX.UserName,
		Token:     tokenX.Token,
	}, erry
	fmt.Println("test search token:", resultCheckToken.Token)

	//check token before POST
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString := resultCheckToken.Token

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
		return "Post success", errx

	} else {
		makeFilterEXP := s.makeFilterToken("Active")
		tokenUpdate := &entity.Token{
			Status: "Inactive",
		}
		errz := s.repository.Update(makeFilterEXP, tokenUpdate)
		fmt.Println("Re-check err wording", err)
		return "Token is expired", errz
	}
}

func (s *ProductsService) makeFilterToken(filters string) (output map[string]interface{}) {
	output = make(map[string]interface{})
	if len(filters) > 0 {
		output["status"] = filters
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
