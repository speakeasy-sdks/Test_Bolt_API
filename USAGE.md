<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	testboltapi "github.com/speakeasy-sdks/Test_Bolt_API"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/operations"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/models/shared"
	"log"
)

func main() {
	s := testboltapi.New()

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag:           shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddressID: testboltapi.String("D4g3h5tBuVYK9"),
				BillingAddressInput: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testboltapi.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.PaymentMethodCreditCardTypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->