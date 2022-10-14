package types

//import "go.mongodb.org/mongo-driver/bson/primitive"

type AddSiteBody struct {
	URL          string `json:"URL" example:"https://www.gmail.com"`
	SiteName     string `json:"sitename" example:"Gmail"`
	UserName     string `json:"userName" example:"dishas.2305"`
	SitePassword string `json:"sitePassword" example:"Admin@1234"`
	Notes        string `json:"notes" example:""`
}

type SitePayload struct {
	URL          string `json:"URL" example:"https://www.gmail.com"`
	SiteName     string `json:"sitename" example:"Gmail"`
	Folder       string `json:"folder" example:"socialMedia"`
	UserName     string `json:"userName" example:"dishas.2305"`
	SitePassword string `json:"sitePassword" example:"Admin@1234"`
	Notes        string `json:"notes" example:""`
}

type CopyPasswordResponse struct {
	SiteName     string `json:"sitename" example:"Gmail"`
	UserName     string `json:"userName" example:"dishas.2305"`
	SitePassword string `json:"sitePassword" example:"Admin@1234"`
}

type EditSitePayload struct {
	UserName     string `json:"userName" example:"dishas.2305"`
	SitePassword string `json:"sitePassword" example:"Admin@1234"`
}
