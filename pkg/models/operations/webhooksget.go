// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"net/http"
)

type WebhooksGetSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *WebhooksGetSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type WebhooksGetRequest struct {
	// The ID of the webhook whose information to retrieve
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *WebhooksGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type WebhooksGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Webhook information has been retrieved
	Webhook *shared.Webhook
}

func (o *WebhooksGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebhooksGetResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
