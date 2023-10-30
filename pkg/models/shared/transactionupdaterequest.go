// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransactionUpdateRequestOrderCart struct {
	// A shopper-facing identifier corresponding to the order reference associated with this transaction.
	DisplayID *string `json:"display_id,omitempty"`
}

func (o *TransactionUpdateRequestOrderCart) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

type TransactionUpdateRequestOrder struct {
	Cart *TransactionUpdateRequestOrderCart `json:"cart,omitempty"`
}

func (o *TransactionUpdateRequestOrder) GetCart() *TransactionUpdateRequestOrderCart {
	if o == nil {
		return nil
	}
	return o.Cart
}

type TransactionUpdateRequest struct {
	Order *TransactionUpdateRequestOrder `json:"order,omitempty"`
}

func (o *TransactionUpdateRequest) GetOrder() *TransactionUpdateRequestOrder {
	if o == nil {
		return nil
	}
	return o.Order
}