# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [DeleteAddress](#deleteaddress) - Delete an existing address
* [DeletePaymentMethod](#deletepaymentmethod) - Delete an existing payment method
* [Detect](#detect) - Determine the existence of a Bolt account
* [GetDetails](#getdetails) - Retrieve account details

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
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, operations.AccountPaymentMethodDeleteRequest{
        XPublishableKey: "<value>",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Detect

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"context"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"log"
)

func main() {
    s := testboltapi.New()

    ctx := context.Background()
    res, err := s.Account.Detect(ctx, operations.AccountExistsRequest{
        XPublishableKey: "<value>",
        Identifier: shared.Identifier{
            IdentifierType: shared.IdentifierTypeEmail,
            IdentifierValue: "alice@example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
            APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Account.GetDetails(ctx, operations.AccountGetRequest{
        XPublishableKey: "<value>",
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |
