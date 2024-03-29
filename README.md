# github.com/speakeasy-sdks/Test_Bolt_API

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
   <a href="https://github.com/speakeasy-sdks/bolt-php/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/Test_Bolt_API
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New(
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Account](docs/sdks/account/README.md)

* [DeleteAddress](docs/sdks/account/README.md#deleteaddress) - Delete an existing address
* [DeletePaymentMethod](docs/sdks/account/README.md#deletepaymentmethod) - Delete an existing payment method
* [Detect](docs/sdks/account/README.md#detect) - Determine the existence of a Bolt account
* [GetDetails](docs/sdks/account/README.md#getdetails) - Retrieve account details


### [Payments.Guest](docs/sdks/guest/README.md)

* [Initialize](docs/sdks/guest/README.md#initialize) - Initialize a Bolt payment for guest shoppers
* [PerformAction](docs/sdks/guest/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending guest payment
* [Update](docs/sdks/guest/README.md#update) - Update an existing guest payment

### [Payments.LoggedIn](docs/sdks/loggedin/README.md)

* [Initialize](docs/sdks/loggedin/README.md#initialize) - Initialize a Bolt payment for logged in shoppers
* [PerformAction](docs/sdks/loggedin/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment
* [Update](docs/sdks/loggedin/README.md#update) - Update an existing payment

### [OAuth](docs/sdks/oauth/README.md)

* [GetToken](docs/sdks/oauth/README.md#gettoken) - Get OAuth token

### [Testing](docs/sdks/testing/README.md)

* [CreateAccount](docs/sdks/testing/README.md#createaccount) - Create a test account
* [GetCreditCard](docs/sdks/testing/README.md#getcreditcard) - Retrieve a test credit card, including its token
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New(
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{environment}.bolt.com/v3` | `environment` (default is `api-sandbox`) |

#### Example

```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New(
		testboltapi.WithServerIndex(0),
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment testboltapi.ServerEnvironment`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New(
		testboltapi.WithServerURL("https://{environment}.bolt.com/v3"),
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKey`     | apiKey       | API key      |
| `Oauth`      | oauth2       | OAuth2 token |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New(
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New()

	operationSecurity := operations.GuestPaymentsInitializeSecurity{
		APIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Payments.Guest.Initialize(ctx, operations.GuestPaymentsInitializeRequest{
		XPublishableKey: "<value>",
		GuestPaymentInitializeRequest: shared.GuestPaymentInitializeRequest{
			Cart: shared.Cart{
				DisplayID:        testboltapi.String("215614191"),
				OrderDescription: testboltapi.String("Order #1234567890"),
				OrderReference:   "order_100",
				Tax: shared.Amount{
					Currency: shared.CurrencyUsd,
					Units:    900,
				},
				Total: shared.Amount{
					Currency: shared.CurrencyUsd,
					Units:    900,
				},
			},
			PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
				shared.PaymentMethodCreditCard{
					DotTag: shared.PaymentMethodCreditCardTagCreditCard,
					BillingAddress: shared.CreateAddressReferenceSchemasInput(
						shared.SchemasInput{
							DotTag:         shared.SchemasTagExplicit,
							Company:        testboltapi.String("ACME Corporation"),
							CountryCode:    shared.SchemasCountryCodeUs,
							Email:          testboltapi.String("alice@example.com"),
							FirstName:      "Alice",
							LastName:       "Baker",
							Locality:       "San Francisco",
							Phone:          testboltapi.String("+14155550199"),
							PostalCode:     "94105",
							Region:         testboltapi.String("CA"),
							StreetAddress1: "535 Mission St, Ste 1401",
							StreetAddress2: testboltapi.String("c/o Shipping Department"),
						},
					),
					Bin:        "411111",
					Expiration: "2029-03",
					Last4:      "1004",
					Network:    shared.NetworkVisa,
					Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				},
			),
			Profile: shared.ProfileCreationData{
				CreateAccount: true,
				Email:         "alice@example.com",
				FirstName:     "Alice",
				LastName:      "Baker",
				Phone:         testboltapi.String("+14155550199"),
			},
		},
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}
	if res.PaymentResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
