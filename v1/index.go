package route

import (

	// "cities/middleware"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Group) {

	gCustomers := e.Group("/customers")
	CustomersGroup(gCustomers)

	gSites := e.Group("/sites")
	SitesGroup(gSites)
}
