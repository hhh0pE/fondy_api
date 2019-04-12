package fondy_api

import (
	"fmt"

	"github.com/pkg/errors"
)

//easyjson:json
type urlResponseWrapper struct {
	Response URLResultResponse `json:"response"`
}

func (urw urlResponseWrapper) ResponseError() error {
	if urw.Response.ErrorCode != 0 || urw.Response.ErrorMessage != "" {
		return errors.New(fmt.Sprintf("Error #%d: %s", urw.Response.ErrorCode, urw.Response.ErrorMessage))
	}
	return nil
}
func (urw urlResponseWrapper) UnwrappedObject() interface{} {
	return &urw.Response
}

//easyjson:json
type URLResultResponse struct {
	CheckoutURL string   `json:"checkout_url"`
	PaymentID   FondyInt `json:"payment_id"`

	ResponseStatus string `json:"response_status"`
	ErrorMessage   string `json:"error_message"`
	ErrorCode      int    `json:"error_code"`
}
