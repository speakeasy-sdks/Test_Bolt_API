// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodSavedPaymentMethodTag string

const (
	PaymentMethodSavedPaymentMethodTagSavedPaymentMethod PaymentMethodSavedPaymentMethodTag = "saved_payment_method"
)

func (e PaymentMethodSavedPaymentMethodTag) ToPointer() *PaymentMethodSavedPaymentMethodTag {
	return &e
}

func (e *PaymentMethodSavedPaymentMethodTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saved_payment_method":
		*e = PaymentMethodSavedPaymentMethodTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodSavedPaymentMethodTag: %v", v)
	}
}

type PaymentMethodSavedPaymentMethod struct {
	DotTag PaymentMethodSavedPaymentMethodTag `json:".tag"`
	// Payment ID of the saved Bolt Payment method.
	ID string `json:"id"`
}

func (o *PaymentMethodSavedPaymentMethod) GetDotTag() PaymentMethodSavedPaymentMethodTag {
	if o == nil {
		return PaymentMethodSavedPaymentMethodTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodSavedPaymentMethod) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
