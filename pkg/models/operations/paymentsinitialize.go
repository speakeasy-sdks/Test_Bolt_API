// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"net/http"
)

type PaymentsInitializeRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey          string                          `header:"style=simple,explode=false,name=X-Publishable-Key"`
	PaymentInitializeRequest shared.PaymentInitializeRequest `request:"mediaType=application/json"`
}

func (o *PaymentsInitializeRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *PaymentsInitializeRequest) GetPaymentInitializeRequest() shared.PaymentInitializeRequest {
	if o == nil {
		return shared.PaymentInitializeRequest{}
	}
	return o.PaymentInitializeRequest
}

type PaymentsInitializeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The payment was successfully initialized, and was either immediately finalized or is pending
	PaymentResponse interface{}
}

func (o *PaymentsInitializeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PaymentsInitializeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PaymentsInitializeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PaymentsInitializeResponse) GetPaymentResponse() interface{} {
	if o == nil {
		return nil
	}
	return o.PaymentResponse
}
