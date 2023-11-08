// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Network - The credit card's network.
type Network string

const (
	NetworkVisa         Network = "visa"
	NetworkMastercard   Network = "mastercard"
	NetworkAmex         Network = "amex"
	NetworkDiscover     Network = "discover"
	NetworkJcb          Network = "jcb"
	NetworkUnionpay     Network = "unionpay"
	NetworkAlliancedata Network = "alliancedata"
	NetworkCitiplcc     Network = "citiplcc"
)

func (e Network) ToPointer() *Network {
	return &e
}

func (e *Network) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "visa":
		fallthrough
	case "mastercard":
		fallthrough
	case "amex":
		fallthrough
	case "discover":
		fallthrough
	case "jcb":
		fallthrough
	case "unionpay":
		fallthrough
	case "alliancedata":
		fallthrough
	case "citiplcc":
		*e = Network(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Network: %v", v)
	}
}

type CreditCard struct {
	// The expiration date, in YYYY-MM format.
	Expiration string `json:"expiration"`
	// The account number's last four digits.
	Last4 string `json:"last4"`
	// The credit card's network.
	Network Network `json:"network"`
}

func (o *CreditCard) GetExpiration() string {
	if o == nil {
		return ""
	}
	return o.Expiration
}

func (o *CreditCard) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *CreditCard) GetNetwork() Network {
	if o == nil {
		return Network("")
	}
	return o.Network
}
