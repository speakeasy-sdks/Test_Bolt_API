# github.com/speakeasy-sdks/Test_Bolt_API

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
   <a href="https://github.com/speakeasy-sdks/bolt-php/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/Test_Bolt_API
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
			APIKey: "",
			Oauth:  "",
		}),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListingInput: shared.AddressListingInput{
			Company:        testboltapi.String("ACME Corporation"),
			CountryCode:    shared.AddressListingCountryCodeUs,
			Email:          testboltapi.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      testboltapi.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testboltapi.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testboltapi.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testboltapi.String("c/o Shipping Department"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddressListing != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Account](docs/sdks/account/README.md)

* [AddAddress](docs/sdks/account/README.md#addaddress) - Add an address
* [AddPaymentMethod](docs/sdks/account/README.md#addpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [DeleteAddress](docs/sdks/account/README.md#deleteaddress) - Delete an existing address
* [DeletePaymentMethod](docs/sdks/account/README.md#deletepaymentmethod) - Delete an existing payment method
* [Detect](docs/sdks/account/README.md#detect) - Determine the existence of a Bolt account
* [GetDetails](docs/sdks/account/README.md#getdetails) - Retrieve account details
* [UpdateAddress](docs/sdks/account/README.md#updateaddress) - Edit an existing address


### [Payments.Guest](docs/sdks/paymentsguest/README.md)

* [Initialize](docs/sdks/paymentsguest/README.md#initialize) - Initialize a Bolt payment for guest shoppers
* [PerformAction](docs/sdks/paymentsguest/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending guest payment
* [Update](docs/sdks/paymentsguest/README.md#update) - Update an existing guest payment

### [Payments.LoggedIn](docs/sdks/paymentsloggedin/README.md)

* [Initialize](docs/sdks/paymentsloggedin/README.md#initialize) - Initialize a Bolt payment for logged in shoppers
* [PerformAction](docs/sdks/paymentsloggedin/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment
* [Update](docs/sdks/paymentsloggedin/README.md#update) - Update an existing payment

### [Testing](docs/sdks/testing/README.md)

* [CreateAccount](docs/sdks/testing/README.md#createaccount) - Create a test account
* [CreateShipmentTracking](docs/sdks/testing/README.md#createshipmenttracking) - Simulate a shipment tracking update
* [GetCreditCard](docs/sdks/testing/README.md#getcreditcard) - Retrieve a test credit card, including its token

### [Transactions](docs/sdks/transactions/README.md)

* [GetDetails](docs/sdks/transactions/README.md#getdetails) - Retrieve transaction details
* [PerformAction](docs/sdks/transactions/README.md#performaction) - Perform an irreversible action (e.g. capture, refund, void) on a transaction
* [Update](docs/sdks/transactions/README.md#update) - Update certain transaction details
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


## Example

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
			APIKey: "",
			Oauth:  "",
		}),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListingInput: shared.AddressListingInput{
			Company:        testboltapi.String("ACME Corporation"),
			CountryCode:    shared.AddressListingCountryCodeUs,
			Email:          testboltapi.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      testboltapi.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testboltapi.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testboltapi.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testboltapi.String("c/o Shipping Department"),
		},
	})
	if err != nil {

		var e *accountAddressCreate_4XXApplicationJSON_OneOf
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{environment}.bolt.com/v3` | `environment` (default is `api-sandbox`) |


Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment ServerEnvironment`

For example:


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
			APIKey: "",
			Oauth:  "",
		}),
		testboltapi.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListingInput: shared.AddressListingInput{
			Company:        testboltapi.String("ACME Corporation"),
			CountryCode:    shared.AddressListingCountryCodeUs,
			Email:          testboltapi.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      testboltapi.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testboltapi.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testboltapi.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testboltapi.String("c/o Shipping Department"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddressListing != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

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
		testboltapi.WithSecurity(shared.Security{
			APIKey: "",
			Oauth:  "",
		}),
		testboltapi.WithServerURL("https://{environment}.bolt.com/v3"),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListingInput: shared.AddressListingInput{
			Company:        testboltapi.String("ACME Corporation"),
			CountryCode:    shared.AddressListingCountryCodeUs,
			Email:          testboltapi.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      testboltapi.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testboltapi.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testboltapi.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testboltapi.String("c/o Shipping Department"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddressListing != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

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
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
