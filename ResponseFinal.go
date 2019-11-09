package fondy_api

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

//easyjson:json
type finalResponseWrapper struct {
	Response FinalResponse `json:"response"`
}

func (urw finalResponseWrapper) ResponseError() error {
	if urw.Response.ErrorCode != 0 || urw.Response.ErrorMessage != "" {
		return errors.New(fmt.Sprintf("Error #%d: %s", urw.Response.ErrorCode, urw.Response.ErrorMessage))
	}
	return nil
}
func (urw finalResponseWrapper) UnwrappedObject() interface{} {
	return &urw.Response
}

//easyjson:json
type FinalResponse struct {
	OrderID     string     `json:"order_id"`
	MerchantID  FondyInt   `json:"merchant_id"`
	Amount      FondyFloat `json:"amount"`
	Currency    string     `json:"currency"`
	OrderStatus string     `json:"order_status"`

	Signature       string `json:"signature"`
	TransactionType string `json:"tran_type"`

	SenderCellPhone string `json:"sender_cell_phone"`
	SenderAccount   string `json:"sender_account"`

	MaskedCard string `json:"masked_card"`

	CardBin  FondyInt `json:"card_bin"`
	CardType string   `json:"card_type"`

	RRN                 string     `json:"rrn"`
	ApprovalCode        string     `json:"approval_code"`
	ResponseCode        FondyInt   `json:"response_code"`
	ResponseDescription string     `json:"response_description"`
	ReversalAmount      FondyFloat `json:"reversal_amount"`

	SettlementAmount FondyFloat `json:"settlement_amount"`

	SettlementCurrency string `json:"settlement_currency"`

	OrderTime FondyTime `json:"order_time"`

	SettlementDate string `json:"settlement_date"`

	Eci string     `json:"eci"`
	Fee FondyFloat `json:"fee"`

	PaymentSystem string `json:"payment_system"`

	SenderEmail string `json:"sender_email"`

	PaymentID FondyInt `json:"payment_id"`

	MerchantData string `json:"merchant_data"`

	VerificationStatus string `json:"verification_status"`

	Rectoken         string    `json:"rectoken"`
	RectokenLifetime FondyTime `json:"rectoken_lifetime"`

	ResponseStatus string   `json:"response_status"`
	ErrorMessage   string   `json:"error_message"`
	ErrorCode      FondyInt `json:"error_code"`

	AdditionalInfo json.RawMessage `json:"additional_info"`
}

func (uad *FinalResponse) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if uad == nil {
		return nil
	}

	e.AddString("OrderID", uad.OrderID)
	e.AddInt64("MerchantID", int64(uad.MerchantID))
	e.AddFloat64("Amount", float64(uad.Amount))

	e.AddString("Currency", uad.Currency)
	e.AddString("OrderStatus", uad.OrderStatus)
	e.AddString("Signature", uad.Signature)
	e.AddString("TransactionType", uad.TransactionType)
	e.AddString("SenderCellPhone", uad.SenderCellPhone)
	e.AddString("SenderAccount", uad.SenderAccount)
	e.AddString("MaskedCard", uad.MaskedCard)
	e.AddInt64("CardBin", int64(uad.CardBin))
	e.AddString("CardType", uad.CardType)
	e.AddString("RRN", uad.RRN)
	e.AddString("ApprovalCode", uad.ApprovalCode)
	e.AddInt64("ResponseCode", int64(uad.ResponseCode))
	e.AddString("ResponseDescription", uad.ResponseDescription)
	e.AddFloat64("ReversalAmount", float64(uad.ReversalAmount))
	e.AddFloat64("SettlementAmount", float64(uad.SettlementAmount))
	e.AddString("SettlementCurrency", uad.SettlementCurrency)

	e.AddTime("OrderTime", uad.OrderTime.Time)
	e.AddString("SettlementDate", uad.SettlementDate)
	e.AddString("Eci", uad.Eci)

	e.AddFloat64("Fee", float64(uad.Fee))

	e.AddString("PaymentSystem", uad.PaymentSystem)

	e.AddString("SenderEmail", uad.SenderEmail)

	e.AddInt64("PaymentID", int64(uad.PaymentID))

	e.AddString("MerchantData", uad.MerchantData)
	e.AddString("VerificationStatus", uad.VerificationStatus)
	e.AddString("Rectoken", uad.Rectoken)
	e.AddString("ResponseStatus", uad.ResponseStatus)
	e.AddString("ErrorMessage", uad.ErrorMessage)

	e.AddTime("RectokenLifetime", uad.RectokenLifetime.Time)

	e.AddInt64("ErrorCode", int64(uad.ErrorCode))

	return nil
}
