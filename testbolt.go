// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package testbolt

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.{username}.dev.bolt.me/v3",
	"https://{environment}.bolt.com/v3",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// TestBolt - Bolt API Reference: A comprehensive Bolt API reference for interacting with Transactions, Orders, Product Catalog, Configuration, Testing, and much more.
type TestBolt struct {
	// Account - Account endpoints allow you to view and manage shoppers' accounts. For example,
	// you can add or remove addresses and payment information.
	//
	Account *account
	// Configuration - Merchant configuration endpoints allow you to retrieve and configure merchant-level
	// configuration, such as callback URLs, identifiers, secrets, etc...
	//
	Configuration *configuration
	// Payments - Use the Payments API to tokenize and process alternative payment methods including Paypal with Bolt. This API is for the Bolt
	// Accounts package.
	//
	Payments *payments
	// Testing - Endpoints that allow you to generate and retrieve test data to verify certain
	// flows in non-production environments.
	//
	Testing *testing
	// Webhooks - Set up webhooks to notify your backend of events within Bolt. These webhooks
	// can communicate with your OMS or other systems to keep them up to date with Bolt.
	//
	// https://help.bolt.com/get-started/during-checkout/webhooks/
	Webhooks *webhooks

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*TestBolt)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *TestBolt) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *TestBolt) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *TestBolt) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

type ServerEnvironment string

const (
	ServerEnvironmentAPI        ServerEnvironment = "api"
	ServerEnvironmentAPISandbox ServerEnvironment = "api-sandbox"
	ServerEnvironmentAPIStaging ServerEnvironment = "api-staging"
)

func (e ServerEnvironment) ToPointer() *ServerEnvironment {
	return &e
}

func (e *ServerEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api":
		fallthrough
	case "api-sandbox":
		fallthrough
	case "api-staging":
		*e = ServerEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerEnvironment: %v", v)
	}
}

// WithEnvironment allows setting the $name variable for url substitution
func WithEnvironment(environment ServerEnvironment) SDKOption {
	return func(sdk *TestBolt) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["environment"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["environment"] = fmt.Sprintf("%v", environment)
		}
	}
}

// WithUsername allows setting the $name variable for url substitution
func WithUsername(username string) SDKOption {
	return func(sdk *TestBolt) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["username"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["username"] = fmt.Sprintf("%v", username)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *TestBolt) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *TestBolt {
	sdk := &TestBolt{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.1",
			SDKVersion:        "0.2.0",
			GenVersion:        "2.89.1",
			ServerDefaults: []map[string]string{
				{
					"username": "xwang",
				},
				{
					"environment": "api-sandbox",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.Account = newAccount(sdk.sdkConfiguration)

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Testing = newTesting(sdk.sdkConfiguration)

	sdk.Webhooks = newWebhooks(sdk.sdkConfiguration)

	return sdk
}
