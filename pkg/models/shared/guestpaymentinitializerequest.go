// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GuestPaymentInitializeRequest struct {
	Cart          Cart                `json:"cart"`
	PaymentMethod PaymentMethod       `json:"payment_method"`
	Profile       ProfileCreationData `json:"profile"`
}

func (o *GuestPaymentInitializeRequest) GetCart() Cart {
	if o == nil {
		return Cart{}
	}
	return o.Cart
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethod() PaymentMethod {
	if o == nil {
		return PaymentMethod{}
	}
	return o.PaymentMethod
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodCreditCard() *PaymentMethodCreditCard {
	return o.GetPaymentMethod().PaymentMethodCreditCard
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodPaypal() *PaymentMethodPaypal {
	return o.GetPaymentMethod().PaymentMethodPaypal
}

func (o *GuestPaymentInitializeRequest) GetProfile() ProfileCreationData {
	if o == nil {
		return ProfileCreationData{}
	}
	return o.Profile
}
