package services

import (
	"passmanager/config"
	"passmanager/models"
	"passmanager/storage"
	"passmanager/types"
	"passmanager/utils"

	"context"
	"fmt"
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	"strconv"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomersReceiver struct {
	MDB           *mongo.Database
	CustomerId    int
	MemberPayload types.CustomerPayload
}

func CreateUser(c *types.CustomerPayload) (interface{}, error) {
	fmt.Println("inside services user")
	um := models.CustomersModel{}
	um.Phone = c.Phone
	hashMPin, err := utils.HashPassword(c.MPin)
	if err != nil {
		logger.Error("func_CreateUser: Error in encrypt password: ", err)
		return nil, err
	}
	um.MPin = hashMPin
	mdb := storage.MONGO_DB
	_, err = mdb.Collection(models.CustomersCollection).InsertOne(context.TODO(), um)
	if err != nil {
		logger.Error("func_CreateUser: ", err)
		return nil, err
	}

	return nil, nil
}

func GetUserByMobileNumber(phone string) (models.CustomersModel, error) {
	var user models.CustomersModel
	mdb := storage.MONGO_DB

	filter := bson.M{
		"phone": phone,
	}

	result := mdb.Collection(models.CustomersCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&user)
	if err != nil {
		logger.Error("func_S_GetGrant: Error in ", err)
		return user, err
	}
	return user, nil
}

func Login(payload types.LoginBody) (types.LoginOutput, error) {
	var loginOutput types.LoginOutput
	user, err := GetUserByMobileNumber(payload.Phone)
	if err != nil {
		logger.Error("GetUserByMobileNumber: Error in fetching customer by mobile number. Error: ", err)
		return loginOutput, err
	}
	if user.Phone == "" {
		return loginOutput, config.ErrUserDoesNotExist
	}
	// reqMPin := payload.Mpin
	// decMPin, err := utils.Decrypt(user.MPin, os.Getenv("MPIN_ENC_KEY"))
	// if reqMPin == decMPin {
	// token, err := GenerateToken(user)
	// if err != nil {
	// 	logger.Error("GenerateToken: Error in generating the token Error: ", err)
	// 	return loginOutput, err
	// }
	// loginOutput.Phone = user.Phone
	// loginOutput.Id = user.ID
	// loginOutput.Token = token
	//

	// 	return loginOutput, nil
	// } else {
	// 	return loginOutput, config.ErrInvalidMPin
	// }
	hashedMpin, err := utils.HashPassword(payload.Mpin)

	if utils.CheckMpinMatch(hashedMpin, payload.Mpin) {
		token, err := GenerateToken(user)
		if err != nil {
			logger.Error("GenerateToken: Error in generating the token Error: ", err)
			return loginOutput, err
		}
		loginOutput.Phone = user.Phone
		loginOutput.Id = user.ID
		loginOutput.Token = token
	}
	return loginOutput, err
}

func ForgotPassword(payload types.ForgotPasswordBody) (types.ForgotPasswordResponse, error) {
	fmt.Println("inside services")
	var ForgotPasswordResponse types.ForgotPasswordResponse
	um := models.CustomersModel{}
	user, err := GetUserByMobileNumber(payload.Phone)
	if err != nil {
		logger.Error("GetUserByMobileNumber: Error in fetching customer by mobile number. Error: ", err)
		return ForgotPasswordResponse, err
	}
	if user.Phone == "" {
		return ForgotPasswordResponse, config.ErrUserDoesNotExist
	}
	encOTP, OTP, err := utils.GenOTP(payload.Phone)
	ForgotPasswordResponse.Phone = user.Phone
	ForgotPasswordResponse.OTP = OTP
	um.OTP = encOTP
	mdb := storage.MONGO_DB
	filter := bson.M{
		"phone": user.Phone,
	}
	update := bson.M{"$set": bson.M{"otp": encOTP}}

	_, err = mdb.Collection(models.CustomersCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logger.Error("func_CreateUser: ", err)
		return ForgotPasswordResponse, err
	}

	return ForgotPasswordResponse, nil

}

func ValidateOTP(user string, otp string) (types.ResetPasswordResponse, error) {
	fmt.Println("inside validate otp")
	var ResetPasswordResponse types.ResetPasswordResponse
	um := models.CustomersModel{}
	decOTP, err := utils.Decrypt(um.OTP, os.Getenv("OTP_ENC_KEY"))
	if err != nil {
		logger.Error("GenerateToken: Error in generating the token Error: ", err)
		return ResetPasswordResponse, err
	}
	if otp == decOTP {
		token, err := GenerateToken(um)
		if err != nil {
			logger.Error("GenerateToken: Error in generating the token Error: ", err)
			return ResetPasswordResponse, err
		}
		ResetPasswordResponse.Token = token
		return ResetPasswordResponse, nil
	} else {
		return ResetPasswordResponse, config.ErrInvalidOTP
	}
}

func UpdateMPin(user string, otp string) error {
	encMPin, err := utils.Encrypt(otp, os.Getenv("MPIN_ENC_KEY"))
	if err != nil {
		logger.Error("func_CreateUser: Error in encrypt password: ", err)
		return err
	}
	mdb := storage.MONGO_DB
	uId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		logger.Error("func Update Mpin ", err)

	}
	filter := bson.M{
		"_id": uId,
	}
	update := bson.M{"$set": bson.M{"mpin": encMPin}}

	_, err = mdb.Collection(models.CustomersCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(userResult models.CustomersModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenExpiredBy, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	claims := token.Claims.(jwt.MapClaims)
	claims["phone"] = userResult.Phone
	claims["user_id"] = userResult.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(tokenExpiredBy)).Unix()

	return token.SignedString([]byte(os.Getenv("CUSTOMER_JWT_SECRET_KEY")))
}

func GenerateRefreshToken(userResult models.CustomersModel) (types.RefreshTokenResponse, error) {
	var tokenRefreshResponse types.RefreshTokenResponse
	token := jwt.New(jwt.SigningMethodHS256)
	tokenExpiredBy, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userResult.ID
	claims["is_refresh"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(tokenExpiredBy)).Unix()
	refershToken, err := token.SignedString([]byte(os.Getenv("CUSTOMER_JWT_SECRET_KEY")))
	tokenRefreshResponse.Token = refershToken
	tokenRefreshResponse.Phone = userResult.Phone
	return tokenRefreshResponse, err
}

func StringToNumber(key string) (int, error) {
	nkey, _ := strconv.Atoi(key)
	return nkey, nil
}
