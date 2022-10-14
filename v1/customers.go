package route

import (
	"passmanager/controllers"
	"passmanager/middleware"

	"github.com/labstack/echo/v4"
)

func CustomersGroup(e *echo.Group) {

	e.POST("/signup", controllers.CreateUser)
	e.GET("/login", controllers.Login)
	e.GET("/forgotpassword", controllers.ForgotPassword, middleware.ValidateCustomerToken)
	e.PUT("/resetpassword/:mobileNumber", controllers.ResetPassword, middleware.ValidateCustomerToken)

}
