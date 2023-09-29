// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AccountAddressDeleteSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
	Oauth  string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *AccountAddressDeleteSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AccountAddressDeleteSecurity) GetOauth() string {
	if o == nil {
		return ""
	}
	return o.Oauth
}

type AccountAddressDeleteRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// The ID of the address to delete
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AccountAddressDeleteRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AccountAddressDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AccountAddressDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
