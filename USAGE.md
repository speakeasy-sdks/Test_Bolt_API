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
	res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
		XPublishableKey: "<value>",
		ID:              "D4g3h5tBuVYK9",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->