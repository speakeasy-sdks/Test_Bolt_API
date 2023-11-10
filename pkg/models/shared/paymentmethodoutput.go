// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/utils"
)

type PaymentMethodOutputType string

const (
	PaymentMethodOutputTypeCreditCard PaymentMethodOutputType = "credit_card"
	PaymentMethodOutputTypePaypal     PaymentMethodOutputType = "paypal"
)

type PaymentMethodOutput struct {
	PaymentMethodCreditCardOutput *PaymentMethodCreditCardOutput
	PaymentMethodPaypalOutput     *PaymentMethodPaypalOutput

	Type PaymentMethodOutputType
}

func CreatePaymentMethodOutputCreditCard(creditCard PaymentMethodCreditCardOutput) PaymentMethodOutput {
	typ := PaymentMethodOutputTypeCreditCard
	typStr := PaymentMethodCreditCardTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethodOutput{
		PaymentMethodCreditCardOutput: &creditCard,
		Type:                          typ,
	}
}

func CreatePaymentMethodOutputPaypal(paypal PaymentMethodPaypalOutput) PaymentMethodOutput {
	typ := PaymentMethodOutputTypePaypal
	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethodOutput{
		PaymentMethodPaypalOutput: &paypal,
		Type:                      typ,
	}
}

func (u *PaymentMethodOutput) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "credit_card":
		paymentMethodCreditCardOutput := new(PaymentMethodCreditCardOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodCreditCardOutput, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCreditCardOutput = paymentMethodCreditCardOutput
		u.Type = PaymentMethodOutputTypeCreditCard
		return nil
	case "paypal":
		paymentMethodPaypalOutput := new(PaymentMethodPaypalOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypalOutput, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodPaypalOutput = paymentMethodPaypalOutput
		u.Type = PaymentMethodOutputTypePaypal
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethodOutput) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodCreditCardOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCardOutput, "", true)
	}

	if u.PaymentMethodPaypalOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypalOutput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
