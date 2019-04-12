package fondy_api

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadFinalResponseFromReader(t *testing.T) {
	var respBuffer bytes.Buffer
	fmt.Fprint(&respBuffer, `{"rrn": "", "masked_card": "444455XXXXXX1111", "sender_cell_phone": "", "response_status": "success", "sender_account": "", "fee": "", "rectoken_lifetime": "", "reversal_amount": "0", "settlement_amount": "0", "actual_amount": "", "order_status": "processing", "response_description": "Declined by antifraud", "verification_status": "", "order_time": "11.04.2019 22:50:15", "actual_currency": "", "order_id": "8", "parent_order_id": "", "merchant_data": "", "tran_type": "purchase", "eci": "", "settlement_date": "", "payment_system": "card", "rectoken": "", "approval_code": "", "merchant_id": 1407229, "settlement_currency": "", "payment_id": 141163874, "product_id": "", "currency": "USD", "card_bin": 444455, "response_code": 1051, "card_type": "VISA", "amount": "100", "sender_email": "demo@gmail.com", "signature": "e131b67e3f8c0beeeb2e9e39cfa6f3df5d3c2d2f"}`)

	finalResp, err := ReadFinalResponseFromReader(&respBuffer)
	fmt.Println("response", finalResp)
	fmt.Println("err", err)
}
