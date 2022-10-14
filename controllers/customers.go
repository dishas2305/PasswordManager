package controllers

import (
	"passmanager/config"
	"passmanager/services"

	"passmanager/types"
	"passmanager/utils"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/text/message"
)

func CreateUser(c echo.Context) error {
	fmt.Println("inside controllers create user")
	user := &types.CustomerPayload{}

	if err := c.Bind(user); err != nil {
		logger.Error("func_CreateUser: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}

	if err := utils.ValidateStruct(user); err != nil {
		logger.Error("func_CreateUser: Error in validating request. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMobNum := utils.CheckForNumbers(user.Phone)
	if !validateMobNum {
		logger.Error("func_CreateCustomer: Error :", config.ErrInvalidMobNum)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrInvalidMobNum)
	}

	validateMPin, err := utils.IsMPinValid(user.MPin)
	if err != nil {
		logger.Error("func_CreateUser: Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}
	if !validateMPin {
		logger.Error("func_CreateUser: Error: ", config.ErrInvalidMPin)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrInvalidMPin)
	}

	comp := utils.CheckMPin(user.MPin, user.ReMPin)
	if comp != 0 {
		logger.Error("func_CreateUser: Error: ", config.ErrMPinDoNotMatch)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrMPinDoNotMatch)
	}

	_, err = services.GetUserByMobileNumber(user.Phone)
	if err == nil {
		logger.Error("func_CreateCustomer: Record found:", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(config.ErrDuplicateCustomer), config.ErrDuplicateCustomer)
	}

	_, err = services.CreateUser(user)
	if err != nil {
		logger.Error("func_CreateUser: Error in creating user:", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, config.MsgUserCreated)

}

func Login(c echo.Context) error {
	body := &types.LoginBody{}
	if err := c.Bind(body); err != nil {
		logger.Error("Login: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}

	if err := utils.ValidateStruct(body); err != nil {
		logger.Error("Login: Error in validating request. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := services.Login(*body)
	if err != nil {
		logger.Error("Login: Error in login. Error: ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, result)
}

func ForgotPassword(c echo.Context) error {
	fmt.Println("inside cntrller")
	body := &types.ForgotPasswordBody{}
	if err := c.Bind(body); err != nil {
		logger.Error("Login: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}

	result, err := services.ForgotPassword(*body)
	if err != nil {
		logger.Error("Forgot Password  ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}
	return utils.HttpSuccessResponse(c, http.StatusOK, result)

}

func ResetPassword(c echo.Context) error {
	fmt.Println("inside reset password")
	userId := c.Param("mobileNumber")
	body := &types.ResetPasswordBody{}
	if err := c.Bind(body); err != nil {
		logger.Error("ResetPassword: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}
	_, err := services.GetUserByMobileNumber(userId)
	if err != nil {
		logger.Error("Error in fetching customer ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}
	result, err := services.ValidateOTP(userId, body.OTP)
	if err != nil {
		logger.Error("Reset Password  ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}
	fmt.Println(result)
	comp := utils.CheckMPin(body.NewMPin, body.NewReMPin)
	if comp != 0 {
		logger.Error("func_CreateUser: Error: ", config.ErrMPinDoNotMatch)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrMPinDoNotMatch)
	}
	errr := services.UpdateMPin(userId, body.NewMPin)
	if errr != nil {
		logger.Error("Error updating Mpin: ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}
	return utils.HttpSuccessResponse(c, http.StatusOK, config.MsgMpinChanged)

}
