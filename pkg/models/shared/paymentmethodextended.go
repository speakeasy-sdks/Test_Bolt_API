// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/Test_Bolt_API/pkg/utils"
)

type PaymentMethodExtendedType string

const (
	PaymentMethodExtendedTypeCreditCard PaymentMethodExtendedType = "credit_card"
	PaymentMethodExtendedTypeID         PaymentMethodExtendedType = "id"
	PaymentMethodExtendedTypePaypal     PaymentMethodExtendedType = "paypal"
)

type PaymentMethodExtended struct {
	PaymentMethodReference  *PaymentMethodReference
	PaymentMethodCreditCard *PaymentMethodCreditCard
	PaymentMethodPaypal     *PaymentMethodPaypal

	Type PaymentMethodExtendedType
}

func CreatePaymentMethodExtendedCreditCard(creditCard PaymentMethodCreditCard) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeCreditCard
	typStr := PaymentMethodCreditCardTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodCreditCard: &creditCard,
		Type:                    typ,
	}
}

func CreatePaymentMethodExtendedID(id PaymentMethodReference) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeID
	typStr := PaymentMethodReferenceTag(typ)
	id.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodReference: &id,
		Type:                   typ,
	}
}

func CreatePaymentMethodExtendedPaypal(paypal PaymentMethodPaypal) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypePaypal
	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodPaypal: &paypal,
		Type:                typ,
	}
}

func (u *PaymentMethodExtended) UnmarshalJSON(data []byte) error {

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
		u.Type = PaymentMethodExtendedTypeCreditCard
		return nil
	case "id":
		paymentMethodReference := new(PaymentMethodReference)
		if err := utils.UnmarshalJSON(data, &paymentMethodReference, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodReference = paymentMethodReference
		u.Type = PaymentMethodExtendedTypeID
		return nil
	case "paypal":
		paymentMethodPaypal := new(PaymentMethodPaypal)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypal, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = PaymentMethodExtendedTypePaypal
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethodExtended) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodReference != nil {
		return utils.MarshalJSON(u.PaymentMethodReference, "", true)
	}

	if u.PaymentMethodCreditCard != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCard, "", true)
	}

	if u.PaymentMethodPaypal != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypal, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
