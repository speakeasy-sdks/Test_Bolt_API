# LoggedIn
(*Payments.LoggedIn*)

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for logged in shoppers
* [PerformAction](#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment
* [Update](#update) - Update an existing payment

## Initialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for logged in shoppers.


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
            APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Initialize(ctx, operations.PaymentsInitializeRequest{
        XPublishableKey: "<value>",
        PaymentInitializeRequest: shared.PaymentInitializeRequest{
            Cart: shared.Cart{
                DisplayID: testboltapi.String("215614191"),
                OrderDescription: testboltapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Tax: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
            },
            PaymentMethod: shared.CreatePaymentMethodExtendedPaymentMethodCreditCard(
                    shared.PaymentMethodCreditCard{
                        DotTag: shared.PaymentMethodCreditCardTagCreditCard,
                        BillingAddress: shared.CreateAddressReferenceSchemasInput(
                                shared.SchemasInput{
                                    DotTag: shared.SchemasTagExplicit,
                                    Company: testboltapi.String("ACME Corporation"),
                                    CountryCode: shared.SchemasCountryCodeUs,
                                    Email: testboltapi.String("alice@example.com"),
                                    FirstName: "Alice",
                                    LastName: "Baker",
                                    Locality: "San Francisco",
                                    Phone: testboltapi.String("+14155550199"),
                                    PostalCode: "94105",
                                    Region: testboltapi.String("CA"),
                                    StreetAddress1: "535 Mission St, Ste 1401",
                                    StreetAddress2: testboltapi.String("c/o Shipping Department"),
                                },
                        ),
                        Bin: "411111",
                        Expiration: "2029-03",
                        Last4: "1004",
                        Network: shared.NetworkVisa,
                        Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                    },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PaymentsInitializeRequest](../../pkg/models/operations/paymentsinitializerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PaymentsInitializeResponse](../../pkg/models/operations/paymentsinitializeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PerformAction

Perform an irreversible action on a pending payment, such as finalizing it.


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
            APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.PerformAction(ctx, operations.PaymentsActionRequest{
        XPublishableKey: "<value>",
        ID: "iKv7t5bgt1gg",
        PaymentActionRequest: shared.PaymentActionRequest{
            DotTag: shared.PaymentActionRequestTagFinalize,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PaymentsActionRequest](../../pkg/models/operations/paymentsactionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PaymentsActionResponse](../../pkg/models/operations/paymentsactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a pending payment


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
            APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Update(ctx, operations.PaymentsUpdateRequest{
        XPublishableKey: "<value>",
        ID: "iKv7t5bgt1gg",
        PaymentUpdateRequest: shared.PaymentUpdateRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PaymentsUpdateRequest](../../pkg/models/operations/paymentsupdaterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PaymentsUpdateResponse](../../pkg/models/operations/paymentsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
