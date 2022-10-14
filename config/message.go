package config

import "errors"

const (
	MsgMemberAdded     = "Member has been added"
	MsgMemberDeleted   = "City has been successfully deleted"
	MsgBookDeleted     = "Book has been deleted"
	MsgBookAdded       = "Book has been added"
	MsgFavAdded        = "City has been added to favourites"
	MsgFavRemoved      = "City has been removed from favourites"
	MsgUserAdded       = "User has been added"
	MsgEmailVerified   = "Email has been verified"
	MsgCheckoutClear   = "Book loaned sucessfully"
	MsgUserCreated     = "Congrats!!!Success Signin to access the vault"
	MsgLoginSuccessful = "Logged in successfully"
	MsgOTPSent         = "OTP sent to Mobile Number"
	MsgMpinChanged     = "MPin changed successfully"
	MsgSiteAdded       = "Site added successfully"
	MsgSiteUpdated     = "Site details updated sucessfully"
)

var (
	ErrMissingBasicAuth            = errors.New("Authorization must be required in header")
	ErrWrongPayload                = errors.New("Wrong payload, please try again")
	ErrRecordNotFound              = errors.New("Record not found")
	ErrParameterMissing            = errors.New("Parameter missing")
	ErrTokenMissing                = errors.New("Error token missing")
	ErrInvalidHashKey              = errors.New("Invalid hash key")
	ErrInvalidHttpMethod           = errors.New("Invalid http method")
	ErrHttpCallBadRequest          = errors.New("Bad request")
	ErrHttpCallUnauthorized        = errors.New("Unauthorized")
	ErrHttpCallNotFound            = errors.New("Call not found")
	ErrHttpCallInternalServerError = errors.New("Internal server error")
	ErrWentWrong                   = errors.New("Something went wrong")
	ErrInvalidMobNum               = errors.New("Invalid mobile number")
	ErrInvalidPasswordFormat       = errors.New("Invalid password format")
	ErrDuplicateCustomer           = errors.New("User already exists with this mobile number")
	ErrVerKeyNotFound              = errors.New("verify key not found")
	ErrEmailAlreadyVerified        = errors.New("Email already verified")
	ErrUserDoesNotExist            = errors.New("User does not exist with this email address")
	ErrEmailNotVerified            = errors.New(" Email not verified")
	ErrInvalidToken                = errors.New("Invalid token")
	ErrRefOnly                     = errors.New("Reference Only")
	//ErrWrongPayload  = errors.New("Wrong payload, please try again")
	//ErrInvalidMobNum = errors.New("Invalid Mobile Number")
	ErrInvalidMPin    = errors.New("Invalid MPin")
	ErrMPinDoNotMatch = errors.New("MPins do not match")
	ErrInvalidOTP     = errors.New("Invalid OTP")
	ErrDuplicateSite  = errors.New("Site with this URL already exists in this folder")
	ErrSiteNotFound   = errors.New("Site not found")
	ErrURLKeyNotFound = errors.New("Site URL not found")
	ErrFileNotFound   = errors.New("File not found")
)
