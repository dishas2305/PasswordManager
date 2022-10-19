package middleware

import (
	"business/models"
	"business/storage"
	"context"
	"fmt"
	"net/http"
	"os"
	"passmanager/config"
	"passmanager/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func ValidateCustomerToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := ExtractCustomerTokenID(c)
		if err != nil {
			return utils.HttpErrorResponse(c, http.StatusUnauthorized, config.ErrHttpCallUnauthorized)
		}
		return next(c)
	}
}
func ExtractCustomerTokenID(c echo.Context) (string, error) {
	mdb := storage.MONGO_DB
	tokenString := c.Request().Header.Get("Authorization")
	fmt.Println("tokenString", tokenString)
	if tokenString == "" {
		return "", config.ErrTokenMissing
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error("func_ValidateCustomerToken: Error in jwt token method. Error: ")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("CUSTOMER_JWT_SECRET_KEY")), nil
	})

	fmt.Println("token:", token)
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("claims----->>>", claims)
	fmt.Println("token", token.Valid)
	if ok && token.Valid {
		uid := fmt.Sprintf("%v", claims["user_id"])
		c.Request().Header.Set("userId", uid)
	}
	find := claims["user_id"]
	filter := bson.M{
		"_id": find,
	}
	result := mdb.Collection(models.CustomersCollection).FindOne(context.TODO(), filter)

	fmt.Println("no error")
	return "", nil

}
