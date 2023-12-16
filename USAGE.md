<!-- Start SDK Example Usage [usage] -->
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
	s := testboltapi.New(
		testboltapi.WithSecurity(shared.Security{
			APIKey: testboltapi.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListing: shared.AddressListingInput{
			Company:        testboltapi.String("ACME Corporation"),
			CountryCode:    shared.CountryCodeUs,
			Email:          testboltapi.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      testboltapi.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testboltapi.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testboltapi.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testboltapi.String("c/o Shipping Department"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddressListing != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->