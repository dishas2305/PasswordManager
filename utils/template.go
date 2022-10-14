package utils

import (
	"bytes"
	"html/template"

	logger "github.com/sirupsen/logrus"
)

var (
	OtpSMS                     = `This message is from {{.OtpFor}}. Your verification code to access your account is {{.OTP}}. This code will expire after {{.Validity}} minutes.`
	SecurityTokenBody          = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/" xmlns:rad="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Request" xmlns:rad1="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Security.Request"><soapenv:Header/><soapenv:Body><tem:RetrieveSecurityToken> <tem:RetrieveSecurityTokenRequest><rad:CarrierCodes><rad:CarrierCode><rad:AccessibleCarrierCode>{{.AccessibleCarrierCode}}</rad:AccessibleCarrierCode></rad:CarrierCode></rad:CarrierCodes><rad1:LogonID>{{.LogonID}}</rad1:LogonID><rad1:Password>{{.Password}}</rad1:Password></tem:RetrieveSecurityTokenRequest></tem:RetrieveSecurityToken></soapenv:Body></soapenv:Envelope>`
	TrainRoutesBody            = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/" xmlns:rad="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Request" xmlns:rad1="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Flight.Request"><soapenv:Header/><soapenv:Body><tem:RetrieveAirportRoutes><tem:RetrieveAirportRoutesRequest><rad:SecurityGUID>{{.SecurityGUID}}</rad:SecurityGUID><rad:CarrierCodes><rad:CarrierCode><rad:AccessibleCarrierCode>{{.AccessibleCarrierCode}}</rad:AccessibleCarrierCode></rad:CarrierCode></rad:CarrierCodes><rad:ClientIPAddress/><rad:HistoricUserName/><rad1:LanguageCode>en</rad1:LanguageCode></tem:RetrieveAirportRoutesRequest></tem:RetrieveAirportRoutes></soapenv:Body></soapenv:Envelope>`
	TrainScheduleInfo          = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/" xmlns:rad="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Request" xmlns:rad1="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Flight.Request"><soapenv:Header/><soapenv:Body><tem:GetFlightScheduleInformation><tem:GetFlightScheduleInformationRequest><rad:SecurityGUID>{{.SecurityGUID}}</rad:SecurityGUID><rad:CarrierCodes><rad:CarrierCode><rad:AccessibleCarrierCode>{{.AccessibleCarrierCode}}</rad:AccessibleCarrierCode></rad:CarrierCode></rad:CarrierCodes><rad:ClientIPAddress>?</rad:ClientIPAddress><rad:HistoricUserName>?</rad:HistoricUserName><rad1:SearchType>{{.SearchType}}</rad1:SearchType><rad1:Origin>{{.Origin}}</rad1:Origin><rad1:Destination>{{.Destination}}</rad1:Destination><rad1:FlightNumber>{{.TrainNumber}}</rad1:FlightNumber><rad1:StartSerachDate>{{.StartSearchDate}}</rad1:StartSerachDate><rad1:EndSearchDate>{{.EndSearchDate}}</rad1:EndSearchDate><rad1:IncludeCancelled>{{.IncludeCancelled}}</rad1:IncludeCancelled></tem:GetFlightScheduleInformationRequest></tem:GetFlightScheduleInformation></soapenv:Body></soapenv:Envelope>`
	RetrieveTrainStatusRequest = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/" xmlns:rad="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Request" xmlns:rad1="http://schemas.datacontract.org/2004/07/Radixx.ConnectPoint.Flight.Request"><soapenv:Header/><soapenv:Body><tem:RetrieveFlightStatus_V1><tem:RetrieveFlightStatusRequest><rad:SecurityGUID>{{.SecurityGUID}}</rad:SecurityGUID><rad:CarrierCodes><rad:CarrierCode><rad:AccessibleCarrierCode>{{.AccessibleCarrierCode}}</rad:AccessibleCarrierCode></rad:CarrierCode></rad:CarrierCodes><rad:ClientIPAddress>?</rad:ClientIPAddress><rad:HistoricUserName/><rad1:SearchType>Route</rad1:SearchType><rad1:Origin>{{.Origin}}</rad1:Origin><rad1:Destination>{{.Destination}}</rad1:Destination><rad1:FlightNumber/><rad1:DepartureDate>{{.DepartureDate}}</rad1:DepartureDate></tem:RetrieveFlightStatusRequest></tem:RetrieveFlightStatus_V1></soapenv:Body></soapenv:Envelope>`
)

type VerifySignupEmailTemplate struct {
	VerifySignupLink  string `validate:"required"`
	PrivacyPolicyLink string `validate:"required"`
	TOSLink           string `validate:"required"`
	UnsubscribeLink   string `validate:"required"`
}

type AccountCreatedTemplate struct {
	BookRideLink      string `validate:"required"`
	PrivacyPolicyLink string `validate:"required"`
	TOSLink           string `validate:"required"`
	CopyRightYear     string `validate:"required"`
	UnsubscribeLink   string `validate:"required"`
}

type ForgotPasswordTemplate struct {
	ForgotPasswordLink          string `validate:"required"`
	ForgotPasswordLinkExpiredBy int    `validate:"required"`
	PrivacyPolicyLink           string `validate:"required"`
	TOSLink                     string `validate:"required"`
	CopyRightYear               string `validate:"required"`
	UnsubscribeLink             string `validate:"required"`
}

type VerificationCodeTemplate struct {
	OTP                int64  `validate:"required"`
	Validity           string `validate:"required"`
	ChangePasswordLink string `validate:"required"`
	PrivacyPolicyLink  string `validate:"required"`
	TOSLink            string `validate:"required"`
	UnsubscribeLink    string `validate:"required"`
}
type OtpSMSTemplate struct {
	OtpFor   string
	OTP      int64
	Validity string
}

type AccountPaswordUpdateTemplate struct {
	FirstName         string `validate:"required"`
	LastName          string `validate:"required"`
	LoginLink         string `validate:"required"`
	PrivacyPolicyLink string `validate:"required"`
	TOSLink           string `validate:"required"`
	CopyRightYear     string `validate:"required"`
	UnsubscribeLink   string `validate:"required"`
}

type SecurityTokenTemplate struct {
	AccessibleCarrierCode string
	LogonID               string
	Password              string
}

type AirportRoutesTemplate struct {
	AccessibleCarrierCode string
	SecurityGUID          string
}

type GetFlightScheduleInfoTemplate struct {
	Origin                string
	Destination           string
	AccessibleCarrierCode string
	SecurityGUID          string
	DepartureDate         string
}

func ParseTemplate(templateContent string, data interface{}) (string, error) {
	t := template.New("NewTemplate")
	t, _ = t.Parse(templateContent)
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		logger.Error("ParseTemplate: Error in executeing data: ", err)
		return "", err
	}
	return buf.String(), nil
}
