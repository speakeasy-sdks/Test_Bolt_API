// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/utils"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard PaymentMethodType = "credit_card"
	PaymentMethodTypePaypal     PaymentMethodType = "paypal"
)

type PaymentMethod struct {
	PaymentMethodCreditCard *PaymentMethodCreditCard
	PaymentMethodPaypal     *PaymentMethodPaypal

	Type PaymentMethodType
}

func CreatePaymentMethodCreditCard(creditCard PaymentMethodCreditCard) PaymentMethod {
	typ := PaymentMethodTypeCreditCard

	typStr := PaymentMethodCreditCardTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethod{
		PaymentMethodCreditCard: &creditCard,
		Type:                    typ,
	}
}

func CreatePaymentMethodPaypal(paypal PaymentMethodPaypal) PaymentMethod {
	typ := PaymentMethodTypePaypal

	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethod{
		PaymentMethodPaypal: &paypal,
		Type:                typ,
	}
}

func (u *PaymentMethod) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "credit_card":
		paymentMethodCreditCard := new(PaymentMethodCreditCard)
		if err := utils.UnmarshalJSON(data, &paymentMethodCreditCard, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCreditCard = paymentMethodCreditCard
		u.Type = PaymentMethodTypeCreditCard
		return nil
	case "paypal":
		paymentMethodPaypal := new(PaymentMethodPaypal)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypal, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = PaymentMethodTypePaypal
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethod) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodCreditCard != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCard, "", true)
	}

	if u.PaymentMethodPaypal != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypal, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
