// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransactionActionVoidTag string

const (
	TransactionActionVoidTagVoid TransactionActionVoidTag = "void"
)

func (e TransactionActionVoidTag) ToPointer() *TransactionActionVoidTag {
	return &e
}

func (e *TransactionActionVoidTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "void":
		*e = TransactionActionVoidTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionActionVoidTag: %v", v)
	}
}

type TransactionActionVoid struct {
	DotTag TransactionActionVoidTag `json:".tag"`
}

func (o *TransactionActionVoid) GetDotTag() TransactionActionVoidTag {
	if o == nil {
		return TransactionActionVoidTag("")
	}
	return o.DotTag
}