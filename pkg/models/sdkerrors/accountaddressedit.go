// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// AccountAddressEdit400ApplicationJSON - The address is invalid and cannot be added
type AccountAddressEdit400ApplicationJSON struct {
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddressEdit400ApplicationJSON{}

func (e *AccountAddressEdit400ApplicationJSON) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
