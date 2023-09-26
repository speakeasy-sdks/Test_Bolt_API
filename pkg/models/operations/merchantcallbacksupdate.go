// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"net/http"
)

type MerchantCallbacksUpdateRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string              `header:"style=simple,explode=false,name=X-Publishable-Key"`
	CallbackUrls    shared.CallbackUrls `request:"mediaType=application/json"`
}

func (o *MerchantCallbacksUpdateRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *MerchantCallbacksUpdateRequest) GetCallbackUrls() shared.CallbackUrls {
	if o == nil {
		return shared.CallbackUrls{}
	}
	return o.CallbackUrls
}

type MerchantCallbacksUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Callbacks URLs were successfully updated
	CallbackUrls *shared.CallbackUrls
}

func (o *MerchantCallbacksUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MerchantCallbacksUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MerchantCallbacksUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MerchantCallbacksUpdateResponse) GetCallbackUrls() *shared.CallbackUrls {
	if o == nil {
		return nil
	}
	return o.CallbackUrls
}
