package products

import (
	"context"
	"fmt"
	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/products"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func (s *ProductsService) Update(ctx context.Context, request *model.FitterUpdateProducts) (string, error) {

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

		if CheckRoleUpdate(user_role) == true {
			makeFilter := s.makeFilterProductsUpdate(request)
			productUpdate := &entity.Products{
				Name:   request.Name,
				Detail: request.Detail,
				Qty:    request.QtyUpdate,
			}
			err := s.repository.Update(makeFilter, productUpdate)

			updateReturn, _ := &model.UpdateResponseProducts{
				Name:   productUpdate.Name,
				Detail: productUpdate.Detail,
				Qty:    productUpdate.Qty,
				Id:     int32(productUpdate.ID),
			}, err
			v := Int32toString(updateReturn.Qty)
			x := updateReturn.Name
			z := updateReturn.Detail
			w := "Update success. " + " Name:" + x + " / Detail:" + z + " / New qty = " + v
			return w, err
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

func (s *ProductsService) makeFilterProductsUpdate(filters *model.FitterUpdateProducts) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.Name) > 0 {
		output["name"] = filters.Name
	}
	if len(filters.Detail) > 0 {
		output["detail"] = filters.Detail
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output
}
