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
	"github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
)

func main() {
    s := Test_Bolt_API.New()
    operationSecurity := operations.MerchantCallbacksGetSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Configuration.MerchantCallbacksGet(ctx, operations.MerchantCallbacksGetRequest{
        XPublishableKey: "illum",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.MerchantCallbacksGetRequest](../../models/operations/merchantcallbacksgetrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.MerchantCallbacksGetSecurity](../../models/operations/merchantcallbacksgetsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


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
	"github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
)

func main() {
    s := Test_Bolt_API.New()
    operationSecurity := operations.MerchantCallbacksUpdateSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Configuration.MerchantCallbacksUpdate(ctx, operations.MerchantCallbacksUpdateRequest{
        XPublishableKey: "vel",
        CallbackUrls: shared.CallbackUrls{
            AccountPage: Test_Bolt_API.String("https://www.example.com/account"),
            BaseDomain: Test_Bolt_API.String("https://www.example.com/"),
            ConfirmationRedirect: Test_Bolt_API.String("https://www.example.com/bolt/redirect"),
            CreateOrder: Test_Bolt_API.String("https://www.example.com/bolt/order"),
            Debug: Test_Bolt_API.String("https://www.example.com/bolt/debug"),
            GetAccount: Test_Bolt_API.String("https://www.example.com/bolt/account"),
            MobileAppDomain: Test_Bolt_API.String("https://m.example.com/"),
            OauthLogout: Test_Bolt_API.String("https://www.example.com/bolt/logout"),
            OauthRedirect: Test_Bolt_API.String("https://www.example.com/bolt/oauth"),
            PrivacyPolicy: Test_Bolt_API.String("https://www.example.com/privacy-policy"),
            ProductInfo: Test_Bolt_API.String("https://www.example.com/bolt/product"),
            RemoteAPI: Test_Bolt_API.String("https://www.example.com/bolt/remote-api"),
            Shipping: Test_Bolt_API.String("https://www.example.com/bolt/shipping"),
            SupportPage: Test_Bolt_API.String("https://www.example.com/help"),
            Tax: Test_Bolt_API.String("https://www.example.com/bolt/tax"),
            TermsOfService: Test_Bolt_API.String("https://www.example.com/terms-of-service"),
            UniversalMerchantAPI: Test_Bolt_API.String("https://www.example.com/bolt/merchant-api"),
            UpdateCart: Test_Bolt_API.String("https://www.example.com/bolt/cart"),
            ValidateAdditionalAccountData: Test_Bolt_API.String("https://www.example.com/bolt/validate-account"),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.MerchantCallbacksUpdateRequest](../../models/operations/merchantcallbacksupdaterequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.MerchantCallbacksUpdateSecurity](../../models/operations/merchantcallbacksupdatesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
)

func main() {
    s := Test_Bolt_API.New()
    operationSecurity := operations.MerchantIdentifiersGetSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Configuration.MerchantIdentifiersGet(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Identifiers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.MerchantIdentifiersGetSecurity](../../models/operations/merchantidentifiersgetsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.MerchantIdentifiersGetResponse](../../models/operations/merchantidentifiersgetresponse.md), error**

