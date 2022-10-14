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

func AddSite(c echo.Context) error {
	uId := c.Param("userId")
	fmt.Println(uId)
	// folder := c.Request().Header.Get("folder")
	// fmt.Println(folder)
	input := &types.SitePayload{}
	if err := c.Bind(input); err != nil {
		logger.Error("func_AddSite: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}
	cr := services.SitesReceiver{}
	cr.SitePayload = *input
	fmt.Println("cr.SitePayload ", cr.SitePayload)
	_, err := services.GetSitebyURL(cr.SitePayload.URL)
	if err == nil {
		logger.Error("func_CreateCustomer: Record found:", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(config.ErrDuplicateCustomer), config.ErrDuplicateSite)
	}
	err = cr.AddSite(uId)
	if err != nil {
		logger.Error("func_AddSite:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrHttpCallInternalServerError)
	}
	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgSiteAdded})
}

func ListSites(c echo.Context) error {
	sr := services.SitesReceiver{}
	sites, err := sr.ListSites()
	if err != nil {
		logger.Error("func_ListSites:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, sites)
}

func CopyPassword(c echo.Context) error {
	fmt.Println("insde copy password")
	site := c.Param("sitename")
	_, err := services.GetSiteByName(site)
	if err != nil {
		logger.Error("func_CopyPassword: Record found:", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(config.ErrSiteNotFound), config.ErrSiteNotFound)
	}
	result, err := services.CopyPassword(site)
	if err != nil {
		logger.Error("Error in Copy Password. Error: ", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(err), err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, result)

}

func EditSite(c echo.Context) error {
	site := c.Param("sitename")
	sr := services.EditSitesReceiver{}
	input := &types.EditSitePayload{}
	if err := c.Bind(input); err != nil {
		logger.Error("func_EditSite: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}

	if err := utils.ValidateStruct(input); err != nil {
		logger.Error("func_EditSite: Error in validating request. Error:", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	sr.EditSitePayload = *input
	err := sr.EditSite(site)
	if err != nil {
		logger.Error("func EditSite:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgSiteUpdated})
}
