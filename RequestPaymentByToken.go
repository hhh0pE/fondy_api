package fondy_api

import "github.com/pkg/errors"

//easyjson:json
type PaymentByToken struct {
	OrderID          string     `json:"order_id"`
	OrderDescription string     `json:"order_desc"`
	Amount           FondyFloat `json:"amount"`
	Currency         string     `json:"currency"`

	SenderEmail string `json:"sender_email,omitempty"`

	MerchantData interface{} `json:"merchant_data,omitempty"`

	LanguageCode string `json:"lang,omitempty"`

	Rectoken string `json:"rectoken,omitempty"`

	Lifetime FondySeconds `json:"lifetime,omitempty"`

	ServerCallbackURL string `json:"server_callback_url,omitempty"`

	CVV2 string `json:"cvv2,omitempty"`

	ClientIP string `json:"client_ip,omitempty"`

	response finalResponseWrapper
}

func (r *PaymentByToken) ResponseObject() Response {
	return &r.response
}

func (r PaymentByToken) URLSuffix() string {
	return "recurring/"
}

func (r *PaymentByToken) PrepareData() error {

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
