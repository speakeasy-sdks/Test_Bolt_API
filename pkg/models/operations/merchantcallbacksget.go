// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"net/http"
)

type MerchantCallbacksGetRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *MerchantCallbacksGetRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

type MerchantCallbacksGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Callbacks URLs were successfully retrieved
	CallbackUrls *shared.CallbackUrls
}

func (o *MerchantCallbacksGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MerchantCallbacksGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MerchantCallbacksGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MerchantCallbacksGetResponse) GetCallbackUrls() *shared.CallbackUrls {
	if o == nil {
		return nil
	}
	return o.CallbackUrls
}
