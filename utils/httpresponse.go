package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type SuccessContent struct {
	Code    int         `json:"code" example:"200"`
	Success bool        `json:"success" default:"true"`
	Data    interface{} `json:"data"`
}
type ErrorContent struct {
	Code    int         `json:"code" example:"401"`
	Success bool        `json:"success" default:"false"`
	Message interface{} `json:"message" example:""`
}
type UnauthorizedErrContent struct {
	Code    int    `json:"code" example:"401"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:"Unauthorized, please try again"`
}
type BadRequestContent struct {
	Code    int    `json:"code" example:"400"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:"bad request"`
}
type NotFoundErrorContent struct {
	Code    int    `json:"code" example:"404"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:"Not Found"`
}
type ConflictErrorContent struct {
	Code    int    `json:"code" example:"409"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:"Data already exists"`
}

type TooManyRequestContent struct {
	Code    int    `json:"code" example:"409"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:""`
}
type InternalServerErrorContent struct {
	Code    int    `json:"code" example:"500"`
	Success bool   `json:"success" default:"false"`
	Message string `json:"message" example:"Internal server error, please try again"`
}
type CustomErrorContent struct {
	Code      int         `json:"code"`
	ErrorCode int         `json:"error_code"`
	Success   bool        `json:"success" default:"false"`
	Message   interface{} `json:"message"`
}

func responseExprired() string {
	headerExp, err := strconv.ParseFloat(os.Getenv("API_RESPONSE_HEADER_EXPIRY"), 200)
	if err != nil {
		logger.Error("func_responseExprired: Error in api response header expiry env: %s", err)
	}
	return time.Now().UTC().Add(time.Duration(headerExp) * time.Minute).String()
}

func HttpSuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	content := SuccessContent{}
	content.Code = statusCode
	content.Data = data
	content.Success = true

	c.Response().Header().Set("expiry", responseExprired())
	return c.JSON(statusCode, content)
}

func HttpErrorResponse(c echo.Context, statusCode int, err error) error {
	content := ErrorContent{}
	content.Code = statusCode
	content.Success = false
	content.Message = err.Error()

	c.Response().Header().Set("expiry", responseExprired())
	return c.JSON(statusCode, content)
}
func HttpCustomErrorResponse(c echo.Context, customCode int, statusCode int, err error) error {
	content := CustomErrorContent{}
	content.Code = statusCode
	content.ErrorCode = customCode
	content.Success = false
	content.Message = err.Error()

	c.Response().Header().Set("expiry", responseExprired())
	return c.JSON(statusCode, content)
}
