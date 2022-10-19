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

// AddSite godoc
// @Summary Adds site ....
// @Description Adds Site
// @Tags Users
// @Accept  json
// @Produce json
// @Success 200 {object} utils.SuccessContent{data=config.MsgSiteAdded }
// @Failure 400 {object} utils.ErrorContent
// @Failure 404 {object} utils.ErrorContent
// @Failure 500 {object} utils.ErrorContent
// @Router /customers [post]
// @Security XAccessToken
// @Security CustomerBasicAuth

func AddSite(c echo.Context) error {
	phone := c.Param("phone")
	input := &types.SitePayload{}
	if err := c.Bind(input); err != nil {
		logger.Error("func_AddSite: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}
	cr := services.SitesReceiver{}
	cr.SitePayload = *input
	fmt.Println("cr.SitePayload ", cr.SitePayload)
	_, err := services.GetSitebyURL(cr.SitePayload.URL, phone)
	if err == nil {
		logger.Error("func_CreateCustomer: Record found:", err)
		return utils.HttpErrorResponse(c, utils.GetStatusCode(config.ErrDuplicateSite), config.ErrDuplicateSite)
	}
	err = cr.AddSite(phone)
	if err != nil {
		logger.Error("func_AddSite:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrHttpCallInternalServerError)
	}
	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgSiteAdded})
}

// ListSite godoc
// @Summary lists all  sites ....
// @Description lists all Sites
// @Tags Users
// @Accept  json
// @Produce json
// @Success 200 {object} utils.SuccessContent{data=[]models.SitesModel }
// @Failure 400 {object} utils.ErrorContent
// @Failure 404 {object} utils.ErrorContent
// @Failure 500 {object} utils.ErrorContent
// @Router /customers [get]
// @Security XAccessToken
// @Security CustomerBasicAuth

func ListSites(c echo.Context) error {
	param := c.Param("param")
	sr := services.SitesReceiver{}
	sites, err := sr.ListSites(param)
	if err != nil {
		logger.Error("func_ListSites:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, sites)
}

// CopyPassword godoc
// @Summary Copies password ....
// @Description Copies passwords
// @Tags Users
// @Accept  json
// @Produce json
// @Success 200 {object} utils.SuccessContent{data=types.CopyPasswordResponse }
// @Failure 400 {object} utils.ErrorContent
// @Failure 404 {object} utils.ErrorContent
// @Failure 500 {object} utils.ErrorContent
// @Router /customers [get]
// @Security XAccessToken
// @Security CustomerBasicAuth

func CopyPassword(c echo.Context) error {
	//fmt.Println("insde copy password")
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

// EditSite godoc
// @Summary edits site ....
// @Description edits site
// @Tags Users
// @Accept  json
// @Produce json
// @Success 200 {object} utils.SuccessContent{data=config.MsgSiteUpdated}
// @Failure 400 {object} utils.ErrorContent
// @Failure 404 {object} utils.ErrorContent
// @Failure 500 {object} utils.ErrorContent
// @Router /customers [get]
// @Security XAccessToken
// @Security CustomerBasicAuth

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
