// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodCreditCardTag string

const (
	PaymentMethodCreditCardTagCreditCard PaymentMethodCreditCardTag = "credit_card"
)

func (e PaymentMethodCreditCardTag) ToPointer() *PaymentMethodCreditCardTag {
	return &e
}

func (e *PaymentMethodCreditCardTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "credit_card":
		*e = PaymentMethodCreditCardTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodCreditCardTag: %v", v)
	}
}

// PaymentMethodCreditCardNetwork - The credit card network.
type PaymentMethodCreditCardNetwork string

const (
	PaymentMethodCreditCardNetworkVisa         PaymentMethodCreditCardNetwork = "visa"
	PaymentMethodCreditCardNetworkMastercard   PaymentMethodCreditCardNetwork = "mastercard"
	PaymentMethodCreditCardNetworkAmex         PaymentMethodCreditCardNetwork = "amex"
	PaymentMethodCreditCardNetworkDiscover     PaymentMethodCreditCardNetwork = "discover"
	PaymentMethodCreditCardNetworkJcb          PaymentMethodCreditCardNetwork = "jcb"
	PaymentMethodCreditCardNetworkUnionpay     PaymentMethodCreditCardNetwork = "unionpay"
	PaymentMethodCreditCardNetworkAlliancedata PaymentMethodCreditCardNetwork = "alliancedata"
	PaymentMethodCreditCardNetworkCitiplcc     PaymentMethodCreditCardNetwork = "citiplcc"
)

func (e PaymentMethodCreditCardNetwork) ToPointer() *PaymentMethodCreditCardNetwork {
	return &e
}

func (e *PaymentMethodCreditCardNetwork) UnmarshalJSON(data []byte) error {
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
		*e = PaymentMethodCreditCardNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodCreditCardNetwork: %v", v)
	}
}

// PaymentMethodCreditCardType - Credit card type
type PaymentMethodCreditCardType string

const (
	PaymentMethodCreditCardTypeCredit PaymentMethodCreditCardType = "credit"
	PaymentMethodCreditCardTypePlcc   PaymentMethodCreditCardType = "plcc"
)

func (e PaymentMethodCreditCardType) ToPointer() *PaymentMethodCreditCardType {
	return &e
}

func (e *PaymentMethodCreditCardType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "credit":
		fallthrough
	case "plcc":
		*e = PaymentMethodCreditCardType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodCreditCardType: %v", v)
	}
}

type PaymentMethodCreditCard struct {
	DotTag PaymentMethodCreditCardTag `json:".tag"`
	// The ID of credit card's billing address
	BillingAddressID    *string           `json:"billing_address_id,omitempty"`
	BillingAddressInput *AddressReference `json:"billing_address_input,omitempty"`
	// The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number.
	Bin string `json:"bin"`
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration string  `json:"expiration"`
	ID         *string `json:"id,omitempty"`
	// The last 4 digits of the credit card number.
	Last4 string `json:"last4"`
	// The credit card network.
	Network PaymentMethodCreditCardNetwork `json:"network"`
	// The Bolt token associated to the credit card.
	Token string `json:"token"`
	// Credit card type
	Type PaymentMethodCreditCardType `json:"type"`
}

func (o *PaymentMethodCreditCard) GetDotTag() PaymentMethodCreditCardTag {
	if o == nil {
		return PaymentMethodCreditCardTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodCreditCard) GetBillingAddressID() *string {
	if o == nil {
		return nil
	}
	return o.BillingAddressID
}

func (o *PaymentMethodCreditCard) GetBillingAddressInput() *AddressReference {
	if o == nil {
		return nil
	}
	return o.BillingAddressInput
}

func (o *PaymentMethodCreditCard) GetBillingAddressInputExplicit() *AddressReferenceExplicit {
	if v := o.GetBillingAddressInput(); v != nil {
		return v.AddressReferenceExplicit
	}
	return nil
}

func (o *PaymentMethodCreditCard) GetBillingAddressInputID() *AddressReferenceID {
	if v := o.GetBillingAddressInput(); v != nil {
		return v.AddressReferenceID
	}
	return nil
}

func (o *PaymentMethodCreditCard) GetBin() string {
	if o == nil {
		return ""
	}
	return o.Bin
}

func (o *PaymentMethodCreditCard) GetExpiration() string {
	if o == nil {
		return ""
	}
	return o.Expiration
}

func (o *PaymentMethodCreditCard) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymentMethodCreditCard) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *PaymentMethodCreditCard) GetNetwork() PaymentMethodCreditCardNetwork {
	if o == nil {
		return PaymentMethodCreditCardNetwork("")
	}
	return o.Network
}

func (o *PaymentMethodCreditCard) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *PaymentMethodCreditCard) GetType() PaymentMethodCreditCardType {
	if o == nil {
		return PaymentMethodCreditCardType("")
	}
	return o.Type
}
