package fondy_api

import (
	"github.com/pkg/errors"
)

//easyjson:json
type CapturePrePurchaseRequest struct {
	OrderID  string     `json:"order_id"`
	Amount   FondyFloat `json:"amount"`
	Currency string     `json:"currency"`

	response urlResponseWrapper
}

func (r *CapturePrePurchaseRequest) ResponseObject() Response {
	return &r.response
}

func (r CapturePrePurchaseRequest) URLSuffix() string {
	return "capture/order_id/"
}

func (r *CapturePrePurchaseRequest) PrepareData() error {

	if r.OrderID == "" {
		return errors.New("OrderID is required")
	}
	if r.Amount == 0 {
		return errors.New("Amount is required")
	}

	return nil
}
