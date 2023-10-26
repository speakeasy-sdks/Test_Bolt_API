# Testing
(*Testing*)

## Overview

Endpoints that allow you to generate and retrieve test data to verify certain
flows in non-production environments.


### Available Operations

* [CreateAccount](#createaccount) - Create a test account
* [CreateShipmentTracking](#createshipmenttracking) - Simulate a shipment tracking update
* [GetCreditCard](#getcreditcard) - Retrieve a test credit card, including its token

## CreateAccount

Create a Bolt shopper account for testing purposes.


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
    res, err := s.Testing.CreateAccount(ctx, operations.TestingAccountCreateRequest{
        XPublishableKey: "string",
        AccountTestCreationDataInput: shared.AccountTestCreationDataInput{
            EmailState: shared.AccountTestCreationDataEmailStateUnverified,
            HasAddress: testboltapi.Bool(true),
            IsMigrated: testboltapi.Bool(true),
            PhoneState: shared.AccountTestCreationDataPhoneStateVerified,
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTestCreationData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.TestingAccountCreateRequest](../../models/operations/testingaccountcreaterequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.TestingAccountCreateSecurity](../../models/operations/testingaccountcreatesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.TestingAccountCreateResponse](../../models/operations/testingaccountcreateresponse.md), error**


## CreateShipmentTracking

Simulate a shipment tracking update, such as those that Bolt would ingest from
third-party shipping providers that would allow updating tracking and delivery
information to shipments associated with orders.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/types"
)

func main() {
    s := testboltapi.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Testing.CreateShipmentTracking(ctx, shared.ShipmentTrackingUpdate{
        DeliveryDate: types.MustTimeFromString("2014-08-23:T06:00:00Z"),
        Status: shared.ShipmentTrackingUpdateStatusInTransit,
        TrackingDetails: []shared.ShipmentTrackingUpdateTrackingDetails{
            shared.ShipmentTrackingUpdateTrackingDetails{
                CountryCode: testboltapi.String("US"),
                EventDate: testboltapi.String("2014-08-21:T14:24:00Z"),
                Locality: testboltapi.String("San Francisco"),
                Message: testboltapi.String("Billing information received"),
                PostalCode: testboltapi.String("94105"),
                Region: testboltapi.String("CA"),
                Status: shared.ShipmentTrackingUpdateTrackingDetailsStatusPreTransit.ToPointer(),
            },
        },
        TrackingNumber: "MockBolt-143292",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.ShipmentTrackingUpdate](../../models/shared/shipmenttrackingupdate.md)                                       | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.TestingShipmentTrackingCreateSecurity](../../models/operations/testingshipmenttrackingcreatesecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.TestingShipmentTrackingCreateResponse](../../models/operations/testingshipmenttrackingcreateresponse.md), error**


## GetCreditCard

Retrieve test credit card information. This includes its token, which is
generated against the `4111 1111 1111 1004` test card.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
)

func main() {
    s := testboltapi.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Testing.GetCreditCard(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.TestingCreditCardGetSecurity](../../models/operations/testingcreditcardgetsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.TestingCreditCardGetResponse](../../models/operations/testingcreditcardgetresponse.md), error**

