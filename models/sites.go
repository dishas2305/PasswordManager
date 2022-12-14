package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SitesCollection = "sites"

type SitesModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone        string             `bson:"phone,omitempty" json:"phone,omitempty"`
	URL          string             `bson:"url" json:"url"`
	SiteName     string             `bson:"sitename" json:"sitename"`
	Folder       string             `bson:"folder,omitempty" json:"folder,omitempty"`
	UserName     string             `bson:"userName" json:"userName"`
	SitePassword string             `bson:"sitePassword" json:"sitePassword"`
	Notes        string             `bson:"notes" json:"notes"`
}
