package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerPayload struct {
	Phone  string `json:"phone" example:"8978456532"`
	MPin   string `json:"mpin" example:"1234"`
	ReMPin string `json:"re-mpin" example:"1234"`
}

type LoginBody struct {
	Phone string `json:"phone"`
	Mpin  string `json:"mpin"`
}

type LoginOutput struct {
	Phone        string             `json:"phone"`
	Token        string             `json:"token"`
	Id           primitive.ObjectID `json:"id"`
	RefreshToken string             `json:"refresh_token"`
}

type ForgotPasswordBody struct {
	Phone string `json:"phone" example:"8978456532"`
}

type ForgotPasswordResponse struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

type ResetPasswordBody struct {
	OTP       string `json:"otp" example:"1256"`
	NewMPin   string `json:"MPin" example:"1234"`
	NewReMPin string `json:"reMPin" example:"1234"`
}

type ResetPasswordResponse struct {
	Id    primitive.ObjectID `json:"id"`
	Token string             `json:"token"`
}

type RefreshTokenBody struct {
	Phone string `json:"phone" example:"8978456532"`
}

type RefreshTokenResponse struct {
	Phone string `json:"phone" example:"8978456532"`
	Token string `json:"token"`
}
