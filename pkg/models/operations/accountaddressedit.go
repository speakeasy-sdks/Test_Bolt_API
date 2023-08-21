// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"net/http"
)

type AccountAddressEditSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
	Oauth  string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *AccountAddressEditSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AccountAddressEditSecurity) GetOauth() string {
	if o == nil {
		return ""
	}
	return o.Oauth
}

type AccountAddressEditRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey     string                     `header:"style=simple,explode=false,name=X-Publishable-Key"`
	AddressListingInput shared.AddressListingInput `request:"mediaType=application/json"`
	// The ID of the address to edit
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AccountAddressEditRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressEditRequest) GetAddressListingInput() shared.AddressListingInput {
	if o == nil {
		return shared.AddressListingInput{}
	}
	return o.AddressListingInput
}

func (o *AccountAddressEditRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AccountAddressEditResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The address was successfully edited
	AddressListing *shared.AddressListing
}

func (o *AccountAddressEditResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressEditResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressEditResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddressEditResponse) GetAddressListing() *shared.AddressListing {
	if o == nil {
		return nil
	}
	return o.AddressListing
}
