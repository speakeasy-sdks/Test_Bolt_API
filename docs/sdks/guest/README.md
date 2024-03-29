# Guest
(*Payments.Guest*)

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for guest shoppers
* [PerformAction](#performaction) - Perform an irreversible action (e.g. finalize) on a pending guest payment
* [Update](#update) - Update an existing guest payment

## Initialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for guest shoppers.


### Example Usage

```go
package main

import(
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"context"
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
            PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
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
            Profile: shared.ProfileCreationData{
                CreateAccount: true,
                Email: "alice@example.com",
                FirstName: "Alice",
                LastName: "Baker",
                Phone: testboltapi.String("+14155550199"),
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

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GuestPaymentsInitializeRequest](../../pkg/models/operations/guestpaymentsinitializerequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.GuestPaymentsInitializeSecurity](../../pkg/models/operations/guestpaymentsinitializesecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.GuestPaymentsInitializeResponse](../../pkg/models/operations/guestpaymentsinitializeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PerformAction

Perform an irreversible action on a pending guest payment, such as finalizing it.


### Example Usage

```go
package main

import(
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
    s := testboltapi.New()


    operationSecurity := operations.GuestPaymentsActionSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Payments.Guest.PerformAction(ctx, operations.GuestPaymentsActionRequest{
        XPublishableKey: "<value>",
        ID: "iKv7t5bgt1gg",
        PaymentActionRequest: shared.PaymentActionRequest{
            DotTag: shared.PaymentActionRequestTagFinalize,
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GuestPaymentsActionRequest](../../pkg/models/operations/guestpaymentsactionrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GuestPaymentsActionSecurity](../../pkg/models/operations/guestpaymentsactionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.GuestPaymentsActionResponse](../../pkg/models/operations/guestpaymentsactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a pending guest payment


### Example Usage

```go
package main

import(
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
    s := testboltapi.New()


    operationSecurity := operations.GuestPaymentsUpdateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Payments.Guest.Update(ctx, operations.GuestPaymentsUpdateRequest{
        XPublishableKey: "<value>",
        ID: "iKv7t5bgt1gg",
        PaymentUpdateRequest: shared.PaymentUpdateRequest{},
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GuestPaymentsUpdateRequest](../../pkg/models/operations/guestpaymentsupdaterequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GuestPaymentsUpdateSecurity](../../pkg/models/operations/guestpaymentsupdatesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.GuestPaymentsUpdateResponse](../../pkg/models/operations/guestpaymentsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
