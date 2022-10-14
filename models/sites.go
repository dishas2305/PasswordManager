package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SitesCollection = "sites"

type SitesModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId       string             `bson:"userId,omitempty" json:"userId,omitempty"`
	URL          string             `bson:"url" json:"url"`
	SiteName     string             `bson:"sitename" json:"sitename"`
	Folder       string             `bson:"folder,omitempty" json:"folder,omitempty"`
	UserName     string             `bson:"userName" json:"userName"`
	SitePassword string             `bson:"sitePassword" json:"sitePassword"`
	Notes        string             `bson:"	notes" json:"notes"`
	ImageResId   string             `bson:"	imageResId" json:"imageResId"`
}
