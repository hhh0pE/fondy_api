package fondy_api

import (
	"github.com/pkg/errors"
)

//easyjson:json
type CreatePaymentRequest struct {
	OrderID          string     `json:"order_id"`
	OrderDescription string     `json:"order_desc"`
	Amount           FondyFloat `json:"amount"`
	Currency         string     `json:"currency"`

	SenderEmail string `json:"sender_email,omitempty"`

	MerchantData interface{} `json:"merchant_data,omitempty"`

	LanguageCode string `json:"lang,omitempty"`

	RequiredRectoken FondyBool `json:"required_rectoken,omitempty"`

	Rectoken string `json:"rectoken,omitempty"`

	Verification     FondyBool `json:"verification,omitempty"`
	VerificationType string    `json:"verification_type,omitempty"`

	Lifetime FondySeconds `json:"lifetime,omitempty"`

	PreAuth FondyBool `json:"preauth"`

	ResponseURL       string `json:"response_url,omitempty"`
	ServerCallbackURL string `json:"server_callback_url,omitempty"`

	response urlResponseWrapper
}

func (r *CreatePaymentRequest) ResponseObject() Response {
	return &r.response
}

func (r CreatePaymentRequest) URLSuffix() string {
	return "checkout/url/"
}

func (r *CreatePaymentRequest) PrepareData() error {

	if r.OrderID == "" {
		return errors.New("OrderID is required")
	}
	if r.OrderDescription == "" {
		return errors.New("OrderDescription is required")
	}
	if r.Amount == 0 {
		return errors.New("Amount is required")
	}

	return nil
}
