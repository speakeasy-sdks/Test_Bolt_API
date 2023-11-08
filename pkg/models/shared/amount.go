// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Currency - A supported currency.
type Currency string

const (
	CurrencyAud Currency = "AUD"
	CurrencyCad Currency = "CAD"
	CurrencyEur Currency = "EUR"
	CurrencyGbp Currency = "GBP"
	CurrencyUsd Currency = "USD"
)

func (e Currency) ToPointer() *Currency {
	return &e
}

func (e *Currency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUD":
		fallthrough
	case "CAD":
		fallthrough
	case "EUR":
		fallthrough
	case "GBP":
		fallthrough
	case "USD":
		*e = Currency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Currency: %v", v)
	}
}

// Amount - A monetary amount, i.e. a base unit amount and a supported currency.
type Amount struct {
	// A supported currency.
	Currency Currency `json:"currency"`
	// A monetary amount, represented in its base units (e.g. USD/EUR cents).
	Units int64 `json:"units"`
}

func (o *Amount) GetCurrency() Currency {
	if o == nil {
		return Currency("")
	}
	return o.Currency
}

func (o *Amount) GetUnits() int64 {
	if o == nil {
		return 0
	}
	return o.Units
}
