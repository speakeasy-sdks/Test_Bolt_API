# Testing
(*Testing*)

## Overview

Endpoints that allow you to generate and retrieve test data to verify certain
flows in non-production environments.


### Available Operations

* [CreateAccount](#createaccount) - Create a test account
* [GetCreditCard](#getcreditcard) - Retrieve a test credit card, including its token

## CreateAccount

Create a Bolt shopper account for testing purposes.


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


    operationSecurity := operations.TestingAccountCreateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Testing.CreateAccount(ctx, operations.TestingAccountCreateRequest{
        XPublishableKey: "string",
        AccountTestCreationData: shared.AccountTestCreationData{
            EmailState: shared.EmailStateUnverified,
            HasAddress: testboltapi.Bool(true),
            IsMigrated: testboltapi.Bool(true),
            PhoneState: shared.PhoneStateVerified,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.TestingAccountCreateRequest](../../pkg/models/operations/testingaccountcreaterequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.TestingAccountCreateSecurity](../../pkg/models/operations/testingaccountcreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.TestingAccountCreateResponse](../../pkg/models/operations/testingaccountcreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## GetCreditCard

Retrieve test credit card information. This includes its token, which is
generated against the `4111 1111 1111 1004` test card.


### Example Usage

```go
package main

import(
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := testboltapi.New()


    operationSecurity := operations.TestingCreditCardGetSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Testing.GetCreditCard(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TestCreditCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.TestingCreditCardGetSecurity](../../pkg/models/operations/testingcreditcardgetsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.TestingCreditCardGetResponse](../../pkg/models/operations/testingcreditcardgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
