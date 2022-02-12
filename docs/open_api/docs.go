// Package classification GoChat API.
//
// the purpose of this service is to provide Authentication related APIs and Create agora tokens
// for authenticated users to join a room related APIs
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /api
//     Version: 1.1.0
//     License: Mozilla Public License
//     Contact: Rezwanul-Haque<rezwanul.haque@vivasoftltd.com> https://www.linkedin.com/in/rezwanul-haque/
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package openapi

import (
	"gochat/app/serializers"
	"gochat/infra/errors"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Payload for singup new user
// swagger:parameters userCreateRequest
type userCreateRequestWrapper struct {
	// in:body
	Body serializers.UserReq
}

// Payload for a user to login
// swagger:parameters loginRequest
type loginRequestWrapper struct {
	// in:body
	Body serializers.UserReq
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

// Payload for renewing a access token using refresh token
// swagger:parameters refreshTokenRequest
type refreshTokenRequestWrapper struct {
	// in:body
	RefreshToken RefreshTokenReq `json:"refresh_token"`
}

// Payload for generating a new agora token
// swagger:parameters createTokenRequest
type createTokenRequestWrapper struct {
	// in:body
	Body serializers.TokenReq
}

// Generic error message
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body errors.RestErr `json:"error_response"`
}

// Firebase login response
// swagger:response firebaseLoginResponse
type loginResponseWrapper struct {
	IDToken      string `json:"idToken"`      // access_token
	Email        string `json:"email"`        // user email address
	RefreshToken string `json:"refreshToken"` // token to renew id token
	ExpiresIn    string `json:"expiresIn"`    // access token will expires in
	LocalID      string `json:"localId"`      // user id
}

// Firebase renew refresh token response
// swagger:response firebaseRenewRefreshResponse
type renewRefreshTokenResponseWapper struct {
	ExpiresIn    string `json:"expires_in"`    // access token will expires in
	TokenType    string `json:"token_type"`    // token type
	RefreshToken string `json:"refresh_token"` // token to renew id token
	IDToken      string `json:"id_token"`      // access_token
	UserID       string `json:"user_id"`       // user id
}

// RTC Token response
// swagger:response rtcTokenResponse
type rtcTokenResponseWrapper struct {
	// rtc response token
	// in: body
	Body serializers.TokenResp `json:"rtc_token"`
}

// returns a 8 digit unique room id in the response
// swagger:response roomResponse
type roomResponseWrapper struct {
	// 8 digit unique room id
	// in: body
	Body serializers.RoomResp `json:"room_id"`
}

type genericSuccessResponse struct {
	Message string `json:"message"`
}

// returns a message
// swagger:response genericSuccessResponse
type genericSuccessResponseWrapper struct {
	// in: body
	genericSuccessResponse `json:"message"`
}

// returns a boolean flag to indicate if the app is online or not
// swagger:response appStatusResponse
type appStatusResponseWrapper struct {
	// a boolean flag to indicate if the app is online or not
	// in: body
	Body serializers.HealthResp `json:"app_online"`
}
