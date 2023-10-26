# PaymentsGuest
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
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
)

func main() {
    s := testboltapi.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Payments.Guest.Initialize(ctx, operations.GuestPaymentsInitializeRequest{
        XPublishableKey: "string",
        GuestPaymentInitializeRequest: shared.GuestPaymentInitializeRequest{
            Cart: shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
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
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testboltapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceExplicitInput(
                                shared.AddressReferenceAddressReferenceExplicitInput{
                                    DotTag: shared.AddressReferenceAddressReferenceExplicitTagExplicit,
                                    Company: testboltapi.String("ACME Corporation"),
                                    CountryCode: shared.AddressReferenceAddressReferenceExplicitCountryCodeUs,
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
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
            },
            PaymentMethod: shared.CreatePaymentMethodInputPaymentMethodCreditCardInput(
                    shared.PaymentMethodCreditCardInput{
                        DotTag: shared.PaymentMethodCreditCardTagCreditCard,
                        BillingAddress: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceExplicitInput(
                                shared.AddressReferenceAddressReferenceExplicitInput{
                                    DotTag: shared.AddressReferenceAddressReferenceExplicitTagExplicit,
                                    Company: testboltapi.String("ACME Corporation"),
                                    CountryCode: shared.AddressReferenceAddressReferenceExplicitCountryCodeUs,
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
                        Network: shared.PaymentMethodCreditCardNetworkVisa,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GuestPaymentsInitializeRequest](../../models/operations/guestpaymentsinitializerequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GuestPaymentsInitializeSecurity](../../models/operations/guestpaymentsinitializesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.GuestPaymentsInitializeResponse](../../models/operations/guestpaymentsinitializeresponse.md), error**


## PerformAction

Perform an irreversible action on a pending guest payment, such as finalizing it.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
)

func main() {
    s := testboltapi.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Payments.Guest.PerformAction(ctx, operations.GuestPaymentsActionRequest{
        XPublishableKey: "string",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GuestPaymentsActionRequest](../../models/operations/guestpaymentsactionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GuestPaymentsActionSecurity](../../models/operations/guestpaymentsactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GuestPaymentsActionResponse](../../models/operations/guestpaymentsactionresponse.md), error**


## Update

Update a pending guest payment


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
)

func main() {
    s := testboltapi.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Payments.Guest.Update(ctx, operations.GuestPaymentsUpdateRequest{
        XPublishableKey: "string",
        ID: "iKv7t5bgt1gg",
        PaymentUpdateRequest: shared.PaymentUpdateRequest{
            Cart: &shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
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
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testboltapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceID(
                                shared.AddressReferenceAddressReferenceID{
                                    DotTag: shared.AddressReferenceAddressReferenceIDTagID,
                                    ID: "D4g3h5tBuVYK9",
                                },
                        ),
                        Carrier: testboltapi.String("FedEx"),
                        Cost: &shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GuestPaymentsUpdateRequest](../../models/operations/guestpaymentsupdaterequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GuestPaymentsUpdateSecurity](../../models/operations/guestpaymentsupdatesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GuestPaymentsUpdateResponse](../../models/operations/guestpaymentsupdateresponse.md), error**

