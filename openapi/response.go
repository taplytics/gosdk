/*
 * Taplytics Universal API
 *
 * The Taplytics Universal API is an API to quickly use Taplytics features and functionality at edge. This API allows the utilization of the Taplytics experimentation  functionality anywhere in the stack, both client- and server-side.  Each call to the Universal API requires two main parameters: token and user_id. - token is unique to each Taplytics project and can be found under Settings -> Project Settings -> Taplytics SDK Key - user_id is your custom defined user ID.  To be able to utilize the Universal API, please ensure that your Taplytics project is setup for usage. You may contact your Account Manager or support@taplytics.com for any questions.
 *
 * API version: 1.1
 * Contact: support@taplytics.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResponse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
