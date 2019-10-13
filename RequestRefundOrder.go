package fondy_api

import (
	"github.com/pkg/errors"
)

//easyjson:json
type RefundOrderRequest struct {
	OrderID string `json:"order_id"`

	Amount   FondyFloat `json:"amount"`
	Currency string     `json:"currency"`

	response finalResponseWrapper
}

func (r *RefundOrderRequest) ResponseObject() Response {
	return &r.response
}

func (r RefundOrderRequest) URLSuffix() string {
	return "reverse/order_id/"
}

func (r *RefundOrderRequest) PrepareData() error {
	if r.OrderID == "" {
		return errors.New("OrderID is required")
	}

	return nil
}
