// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/utils"
)

type GuestPaymentMethodInitializeRequestPaymentMethodType string

const (
	GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal GuestPaymentMethodInitializeRequestPaymentMethodType = "payment-method-paypal"
)

type GuestPaymentMethodInitializeRequestPaymentMethod struct {
	PaymentMethodPaypal *PaymentMethodPaypal

	Type GuestPaymentMethodInitializeRequestPaymentMethodType
}

func CreateGuestPaymentMethodInitializeRequestPaymentMethodPaymentMethodPaypal(paymentMethodPaypal PaymentMethodPaypal) GuestPaymentMethodInitializeRequestPaymentMethod {
	typ := GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal

	return GuestPaymentMethodInitializeRequestPaymentMethod{
		PaymentMethodPaypal: &paymentMethodPaypal,
		Type:                typ,
	}
}

func (u *GuestPaymentMethodInitializeRequestPaymentMethod) UnmarshalJSON(data []byte) error {

	paymentMethodPaypal := new(PaymentMethodPaypal)
	if err := utils.UnmarshalJSON(data, &paymentMethodPaypal, "", true, true); err == nil {
		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GuestPaymentMethodInitializeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodPaypal != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypal, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GuestPaymentMethodInitializeRequest struct {
	Cart          Cart                                             `json:"cart"`
	PaymentMethod GuestPaymentMethodInitializeRequestPaymentMethod `json:"payment_method"`
}

func (o *GuestPaymentMethodInitializeRequest) GetCart() Cart {
	if o == nil {
		return Cart{}
	}
	return o.Cart
}

func (o *GuestPaymentMethodInitializeRequest) GetPaymentMethod() GuestPaymentMethodInitializeRequestPaymentMethod {
	if o == nil {
		return GuestPaymentMethodInitializeRequestPaymentMethod{}
	}
	return o.PaymentMethod
}
