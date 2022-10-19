package route

import (
	"passmanager/controllers"
	"passmanager/middleware"

	"github.com/labstack/echo/v4"
)

func SitesGroup(e *echo.Group) {
	e.POST("/:phone/addsite", controllers.AddSite, middleware.ValidateCustomerToken)
	e.GET("/:param", controllers.ListSites)
	e.GET("/:sitename/copypassword", controllers.CopyPassword, middleware.ValidateCustomerToken)
	e.PUT("/:sitename", controllers.EditSite, middleware.ValidateCustomerToken)
}
