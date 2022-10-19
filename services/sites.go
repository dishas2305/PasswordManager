package services

import (
	"context"
	"fmt"
	"os"
	"passmanager/models"
	"passmanager/storage"
	"passmanager/types"
	"passmanager/utils"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SitesReceiver struct {
	MDB        *mongo.Database
	CustomerId int

	SitePayload types.SitePayload
}

type EditSitesReceiver struct {
	MDB        *mongo.Database
	CustomerId int

	EditSitePayload types.EditSitePayload
}

func (sr *SitesReceiver) AddSite(userID string) error {
	mdb := storage.MONGO_DB
	_, err := GetUserByMobileNumber(userID)
	if err != nil {
		logger.Error("GetUserByMobileNumber: Error in fetching customer by mobile number. Error: ", err)
		return err
	}
	sm := models.SitesModel{}
	sm.Phone = userID
	sm.URL = sr.SitePayload.URL
	sm.SiteName = sr.SitePayload.SiteName
	sm.Folder = sr.SitePayload.Folder
	sm.UserName = sr.SitePayload.UserName
	encSitePassword, err := utils.Encrypt(sr.SitePayload.SitePassword, os.Getenv("MPIN_ENC_KEY"))
	if err != nil {
		logger.Error("func_CreateUser: Error in encrypt password: ", err)
		return err
	}
	sm.SitePassword = encSitePassword
	sm.Notes = sr.SitePayload.Notes

	_, errr := mdb.Collection(models.SitesCollection).InsertOne(context.TODO(), sm)
	if errr != nil {
		logger.Error("func_AddSite: ", errr)
		return errr
	}
	return nil
}

func GetSitebyURL(siteurl, phone string) (models.SitesModel, error) {
	var site models.SitesModel
	mdb := storage.MONGO_DB
	filter := bson.M{
		"url": siteurl,
	}
	result := mdb.Collection(models.SitesCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&site)
	if err != nil {
		logger.Error("func GetSitebyURL: Error in ", err)
		return site, err
	}
	return site, nil
}

func GetSiteByName(sitename string) (models.SitesModel, error) {
	var site models.SitesModel
	mdb := storage.MONGO_DB
	filter := bson.M{
		"sitename": sitename,
	}
	result := mdb.Collection(models.SitesCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&site)
	if err != nil {
		logger.Error("func GetSitebyURL: Error in ", err)
		return site, err
	}
	return site, nil
}

func (sr *SitesReceiver) ListSites(param string) ([]models.SitesModel, error) {
	mdb := storage.MONGO_DB
	filter := bson.M{
		"folder": param,
	}
	var sites []models.SitesModel
	result, _ := mdb.Collection(models.SitesCollection).Find(context.TODO(), filter)
	fmt.Println(result)
	if err := result.All(context.Background(), &sites); err != nil {
		logger.Error("func_ListSites: error cur.All() step ", err)
		return nil, err
	}

	return sites, nil
}

func CopyPassword(sitename string) (types.CopyPasswordResponse, error) {
	var site models.SitesModel
	var copypassword types.CopyPasswordResponse
	mdb := storage.MONGO_DB
	filter := bson.M{
		"sitename": sitename,
	}
	result := mdb.Collection(models.SitesCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&site)
	if err != nil {
		logger.Error("func GetSitebyURL: Error in ", err)
		return copypassword, err
	}
	decPassword, err := utils.Decrypt(site.SitePassword, os.Getenv("MPIN_ENC_KEY"))
	copypassword.SiteName = site.SiteName
	copypassword.UserName = site.UserName
	copypassword.SitePassword = decPassword

	return copypassword, nil

}

func (sr *EditSitesReceiver) EditSite(sitename string) error {

	mdb := storage.MONGO_DB

	filter := bson.M{
		"sitename": sitename,
	}
	encSitePassword, err := utils.Encrypt(sr.EditSitePayload.SitePassword, os.Getenv("MPIN_ENC_KEY"))
	update := bson.M{"$set": bson.M{"userName": sr.EditSitePayload.UserName, "sitePassword": encSitePassword}}

	_, err = mdb.Collection(models.SitesCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
