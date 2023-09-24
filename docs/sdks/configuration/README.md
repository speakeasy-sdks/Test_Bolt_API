# Configuration

## Overview

Merchant configuration endpoints allow you to retrieve and configure merchant-level
configuration, such as callback URLs, identifiers, secrets, etc...


### Available Operations

* [MerchantCallbacksGet](#merchantcallbacksget) - Retrieve callback URLs for the merchant
* [MerchantCallbacksUpdate](#merchantcallbacksupdate) - Update callback URLs for the merchant
* [MerchantIdentifiersGet](#merchantidentifiersget) - Retrieve identifiers for the merchant

## MerchantCallbacksGet

Return callback URLs configured on the merchant such as OAuth URLs.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.MerchantCallbacksGet(ctx, operations.MerchantCallbacksGetRequest{
        XPublishableKey: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.MerchantCallbacksGetRequest](../../models/operations/merchantcallbacksgetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.MerchantCallbacksGetResponse](../../models/operations/merchantcallbacksgetresponse.md), error**


## MerchantCallbacksUpdate

Update and configure callback URLs on the merchant such as OAuth URLs.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.MerchantCallbacksUpdate(ctx, operations.MerchantCallbacksUpdateRequest{
        XPublishableKey: "error",
        CallbackUrls: shared.CallbackUrls{
            AccountPage: testboltapi.String("https://www.example.com/account"),
            BaseDomain: testboltapi.String("https://www.example.com/"),
            ConfirmationRedirect: testboltapi.String("https://www.example.com/bolt/redirect"),
            CreateOrder: testboltapi.String("https://www.example.com/bolt/order"),
            Debug: testboltapi.String("https://www.example.com/bolt/debug"),
            GetAccount: testboltapi.String("https://www.example.com/bolt/account"),
            MobileAppDomain: testboltapi.String("https://m.example.com/"),
            OauthLogout: testboltapi.String("https://www.example.com/bolt/logout"),
            OauthRedirect: testboltapi.String("https://www.example.com/bolt/oauth"),
            PrivacyPolicy: testboltapi.String("https://www.example.com/privacy-policy"),
            ProductInfo: testboltapi.String("https://www.example.com/bolt/product"),
            RemoteAPI: testboltapi.String("https://www.example.com/bolt/remote-api"),
            Shipping: testboltapi.String("https://www.example.com/bolt/shipping"),
            SupportPage: testboltapi.String("https://www.example.com/help"),
            Tax: testboltapi.String("https://www.example.com/bolt/tax"),
            TermsOfService: testboltapi.String("https://www.example.com/terms-of-service"),
            UniversalMerchantAPI: testboltapi.String("https://www.example.com/bolt/merchant-api"),
            UpdateCart: testboltapi.String("https://www.example.com/bolt/cart"),
            ValidateAdditionalAccountData: testboltapi.String("https://www.example.com/bolt/validate-account"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.MerchantCallbacksUpdateRequest](../../models/operations/merchantcallbacksupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.MerchantCallbacksUpdateResponse](../../models/operations/merchantcallbacksupdateresponse.md), error**


## MerchantIdentifiersGet

Return several identifiers for the merchant, such as public IDs, publishable keys, signing secrets, etc...

### Example Usage

```go
package main

import(
	"context"
	"log"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
)

func main() {
    s := testboltapi.New(
        testboltapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.MerchantIdentifiersGet(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Identifiers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.MerchantIdentifiersGetResponse](../../models/operations/merchantidentifiersgetresponse.md), error**

