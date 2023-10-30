# Transactions
(*Transactions*)

## Overview

Transaction endpoints allow you to manage transactions. For example, you can capture
funds, void transactions, or issue refunds. You can also update certain fields for existing
transactions.


### Available Operations

* [GetDetails](#getdetails) - Retrieve transaction details
* [PerformAction](#performaction) - Perform an irreversible action (e.g. capture, refund, void) on a transaction
* [Update](#update) - Update certain transaction details

## GetDetails

Retrieve information for a specific transaction

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
    res, err := s.Transactions.GetDetails(ctx, operations.TransactionGetRequest{
        XPublishableKey: "string",
        ID: "OBYG-X1PX-FN55",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.TransactionGetRequest](../../models/operations/transactiongetrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.TransactionGetSecurity](../../models/operations/transactiongetsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.TransactionGetResponse](../../models/operations/transactiongetresponse.md), error**


## PerformAction

Perform an irreversible action (e.g. capture, refund, void) on a transaction

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
    res, err := s.Transactions.PerformAction(ctx, operations.TransactionActionRequest{
        XPublishableKey: "string",
        ID: "OBYG-X1PX-FN55",
        TransactionActionRequest: shared.TransactionActionRequest{
            Action: shared.CreateTransactionActionRequestActionTransactionActionCapture(
                    shared.TransactionActionCapture{
                        DotTag: shared.TransactionActionCaptureTagCapture,
                        Amount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                    },
            ),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TransactionActionRequest](../../models/operations/transactionactionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.TransactionActionSecurity](../../models/operations/transactionactionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.TransactionActionResponse](../../models/operations/transactionactionresponse.md), error**


## Update

Update certain transaction details, such as the user-facing ID of its associate order

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
    res, err := s.Transactions.Update(ctx, operations.TransactionUpdateRequest{
        XPublishableKey: "string",
        ID: "OBYG-X1PX-FN55",
        TransactionUpdateRequest: shared.TransactionUpdateRequest{
            Order: &shared.TransactionUpdateRequestOrder{
                Cart: &shared.TransactionUpdateRequestOrderCart{
                    DisplayID: testboltapi.String("215614191"),
                },
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TransactionUpdateRequest](../../models/operations/transactionupdaterequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.TransactionUpdateSecurity](../../models/operations/transactionupdatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.TransactionUpdateResponse](../../models/operations/transactionupdateresponse.md), error**
