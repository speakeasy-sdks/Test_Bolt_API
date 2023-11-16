# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [AddAddress](#addaddress) - Add an address
* [AddPaymentMethod](#addpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [DeleteAddress](#deleteaddress) - Delete an existing address
* [DeletePaymentMethod](#deletepaymentmethod) - Delete an existing payment method
* [Detect](#detect) - Determine the existence of a Bolt account
* [GetDetails](#getdetails) - Retrieve account details
* [UpdateAddress](#updateaddress) - Edit an existing address

## AddAddress

Add an address to the shopper's account

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
        XPublishableKey: "string",
        AddressListing: shared.AddressListingInput{
            Company: testboltapi.String("ACME Corporation"),
            CountryCode: shared.CountryCodeUs,
            Email: testboltapi.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: testboltapi.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: testboltapi.String("+14155550199"),
            PostalCode: "94105",
            Region: testboltapi.String("CA"),
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountAddressCreateRequest](../../pkg/models/operations/accountaddresscreaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.AccountAddressCreateResponse](../../pkg/models/operations/accountaddresscreateresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountAddressCreateResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 400-600                                    | */*                                        |

## AddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from
your backend because authentication requires the use of your private key.<br />
**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function,
which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.AddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
        XPublishableKey: "string",
        PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
                shared.PaymentMethodCreditCard{
                    DotTag: shared.PaymentMethodCreditCardTagCreditCard,
                    BillingAddress: shared.CreateAddressReferenceSchemas(
                            shared.Schemas{
                                DotTag: shared.SchemasAddressReferenceIDTagID,
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Bin: "411111",
                    Expiration: "2029-03",
                    Last4: "1004",
                    Network: shared.PaymentMethodCreditCardNetworkVisa,
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountAddPaymentMethodRequest](../../pkg/models/operations/accountaddpaymentmethodrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.AccountAddPaymentMethodResponse](../../pkg/models/operations/accountaddpaymentmethodresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.AccountAddPaymentMethodResponseBody | 4XX                                           | application/json                              |
| sdkerrors.SDKError                            | 400-600                                       | */*                                           |

## DeleteAddress

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
        XPublishableKey: "string",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountAddressDeleteRequest](../../pkg/models/operations/accountaddressdeleterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.AccountAddressDeleteResponse](../../pkg/models/operations/accountaddressdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## DeletePaymentMethod

Delete an existing payment method. Deleting a payment method does not invalidate transactions or
orders that are associated with it.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, operations.AccountPaymentMethodDeleteRequest{
        XPublishableKey: "string",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountPaymentMethodDeleteRequest](../../pkg/models/operations/accountpaymentmethoddeleterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.AccountPaymentMethodDeleteResponse](../../pkg/models/operations/accountpaymentmethoddeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## Detect

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.Detect(ctx, operations.AccountExistsRequest{
        XPublishableKey: "string",
        Identifier: shared.Identifier{
            IdentifierType: shared.IdentifierTypeEmail,
            IdentifierValue: "alice@example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AccountExistsRequest](../../pkg/models/operations/accountexistsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.AccountExistsResponse](../../pkg/models/operations/accountexistsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## GetDetails

Retrieve a shopper's account details, such as addresses and payment information

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.GetDetails(ctx, operations.AccountGetRequest{
        XPublishableKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AccountGetRequest](../../pkg/models/operations/accountgetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.AccountGetResponse](../../pkg/models/operations/accountgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateAddress

Edit an existing address on the shopper's account. This does not edit addresses
that are already associated with other resources, such as transactions or
shipments.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Account.UpdateAddress(ctx, operations.AccountAddressEditRequest{
        XPublishableKey: "string",
        AddressListing: shared.AddressListingInput{
            Company: testboltapi.String("ACME Corporation"),
            CountryCode: shared.CountryCodeUs,
            Email: testboltapi.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: testboltapi.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: testboltapi.String("+14155550199"),
            PostalCode: "94105",
            Region: testboltapi.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: testboltapi.String("c/o Shipping Department"),
        },
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountAddressEditRequest](../../pkg/models/operations/accountaddresseditrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.AccountAddressEditResponse](../../pkg/models/operations/accountaddresseditresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.AccountAddressEditResponseBody | 4XX                                      | application/json                         |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |
