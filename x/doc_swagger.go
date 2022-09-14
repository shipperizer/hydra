// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

// Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
// typically 201.
//
// swagger:response emptyResponse
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type emptyResponse struct{}

// Error
//
// swagger:model errorOAuth2
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorOAuth2 struct {
	// Error
	Name string `json:"error"`

	// Error Description
	Description string `json:"error_description"`

	// Error Hint
	//
	// Helps the user identify the error cause.
	//
	// Example: The redirect URL is not allowed.
	Hint string `json:"error_hint"`

	// HTTP Status Code
	//
	// Example: 401
	Code int `json:"status_code"`

	// Error Debug Information
	//
	// Only available in dev mode.
	Debug string `json:"error_debug,omitempty"`
}

// Default Error Response
//
// swagger:response errorOAuth2Default
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorOAuth2Default struct {
	// in: body
	Body errorOAuth2
}

// Bad Request Error Response
//
// swagger:response errorOAuth2BadRequest
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorOAuth2BadRequest struct {
	// in: body
	Body errorOAuth2
}

// Not Found Error Response
//
// swagger:response errorOAuth2NotFound
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorOAuth2NotFound struct {
	// in: body
	Body errorOAuth2
}

// OAuth2 Device Flow
//
// # Ory's OAuth 2.0 Device Authorization API
//
// swagger:model oAuth2ApiDeviceAuthorizationResponse
type oAuth2ApiDeviceAuthorizationResponse struct {
	// The device verification code.
	//
	// example: ory_dc_smldfksmdfkl.mslkmlkmlk
	DeviceCode string `json:"device_code"`

	// The end-user verification code.
	//
	// example: AAAAAA
	UserCode string `json:"user_code"`

	// The end-user verification URI on the authorization
	// server.  The URI should be short and easy to remember as end users
	// will be asked to manually type it into their user agent.
	//
	// example: https://auth.ory.sh/tv
	VerificationUri string `json:"verification_uri"`

	// A verification URI that includes the "user_code" (or
	// other information with the same function as the "user_code"),
	// which is designed for non-textual transmission.
	//
	// example: https://auth.ory.sh/tv?user_code=AAAAAA
	VerificationUriComplete string `json:"verification_uri_complete"`

	// The lifetime in seconds of the "device_code" and "user_code".
	//
	// example: 16830
	ExpiresIn int `json:"expires_in"`

	// The minimum amount of time in seconds that the client
	// SHOULD wait between polling requests to the token endpoint.  If no
	// value is provided, clients MUST use 5 as the default.
	//
	// example: 5
	Interval int `json:"interval"`
}
