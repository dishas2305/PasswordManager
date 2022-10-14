package route

import (

	// "cities/middleware"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Group) {
	//e.GET("/health", controllers.HealthCheck)
	//Members Group

	gCustomers := e.Group("/customers")
	CustomersGroup(gCustomers)

	gSites := e.Group("/sites")
	SitesGroup(gSites)
}
