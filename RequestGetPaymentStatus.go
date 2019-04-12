package fondy_api

import (
	"github.com/pkg/errors"
)

//easyjson:json
type GetPaymentStatusRequest struct {
	OrderID string `json:"order_id"`

	response finalResponseWrapper
}

func (r *GetPaymentStatusRequest) ResponseObject() Response {
	return &r.response
}

func (r GetPaymentStatusRequest) URLSuffix() string {
	return "status/order_id/"
}

func (r *GetPaymentStatusRequest) PrepareData() error {

	if r.OrderID == "" {
		return errors.New("OrderID is required")
	}

	return nil
}
