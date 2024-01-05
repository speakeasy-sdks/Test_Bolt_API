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
        XPublishableKey: "string",
        PaymentInitializeRequest: shared.PaymentInitializeRequest{
            Cart: shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                        Code: testboltapi.String("SUMMER10DISCOUNT"),
                        DetailsURL: testboltapi.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: testboltapi.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: testboltapi.String("Large tote with Bolt logo."),
                        ImageURL: testboltapi.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testboltapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceSchemasInput(
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
                        Carrier: testboltapi.String("FedEx"),
                        Cost: &shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
            },
            PaymentMethod: shared.CreatePaymentMethodExtendedPaymentMethodReference(
                    shared.PaymentMethodReference{
                        DotTag: shared.PaymentMethodReferenceTagID,
                        ID: "id",
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
        XPublishableKey: "string",
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
        XPublishableKey: "string",
        ID: "iKv7t5bgt1gg",
        PaymentUpdateRequest: shared.PaymentUpdateRequest{
            Cart: &shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                        Code: testboltapi.String("SUMMER10DISCOUNT"),
                        DetailsURL: testboltapi.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: testboltapi.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: testboltapi.String("Large tote with Bolt logo."),
                        ImageURL: testboltapi.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testboltapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceSchemas(
                                shared.Schemas{
                                    DotTag: shared.SchemasAddressReferenceIDTagID,
                                    ID: "D4g3h5tBuVYK9",
                                },
                        ),
                        Carrier: testboltapi.String("FedEx"),
                        Cost: &shared.Amount{
                            Currency: shared.CurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.CurrencyUsd,
                    Units: 900,
                },
            },
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
| `request`                                                                                | [operations.PaymentsUpdateRequest](../../pkg/models/operations/paymentsupdaterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PaymentsUpdateResponse](../../pkg/models/operations/paymentsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
