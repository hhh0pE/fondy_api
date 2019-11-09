package fondy_api

import (
	"bytes"
	"fmt"
	"net/url"
	"testing"
)

func TestReadFinalResponseFromReader(t *testing.T) {
	var respBuffer bytes.Buffer
	fmt.Fprint(&respBuffer, `{"rrn": "", "masked_card": "444455XXXXXX1111", "sender_cell_phone": "", "response_status": "success", "sender_account": "", "fee": "", "rectoken_lifetime": "", "reversal_amount": "0", "settlement_amount": "0", "actual_amount": "", "order_status": "processing", "response_description": "Declined by antifraud", "verification_status": "", "order_time": "11.04.2019 22:50:15", "actual_currency": "", "order_id": "8", "parent_order_id": "", "merchant_data": "", "tran_type": "purchase", "eci": "", "settlement_date": "", "payment_system": "card", "rectoken": "", "approval_code": "", "merchant_id": 1407229, "settlement_currency": "", "payment_id": 141163874, "product_id": "", "currency": "USD", "card_bin": 444455, "response_code": 1051, "card_type": "VISA", "amount": "100", "sender_email": "demo@gmail.com", "signature": "e131b67e3f8c0beeeb2e9e39cfa6f3df5d3c2d2f"}`)

	finalResp, err := ReadFinalResponseFromReader(&respBuffer)
	fmt.Println("response", finalResp)
	fmt.Println("err", err)
}

func TestParseRequestFormData(t *testing.T) {
	parsedValues, parsing_err := url.ParseQuery("rrn=&masked_card=537541XXXXXX1434&sender_cell_phone=&response_status=success&sender_account=&fee=&rectoken_lifetime=01.07.2024+00%3A00%3A00&reversal_amount=0&settlement_amount=0&actual_amount=2465&order_status=approved&response_description=&verification_status=&order_time=09.11.2019+12%3A01%3A56&actual_currency=UAH&order_id=mbm_challenge40_13&parent_order_id=&merchant_data=&tran_type=purchase&eci=&settlement_date=&payment_system=card&rectoken=36ddc6e83d9ddce7d4dc89fd06b2ffac6cc89982&approval_code=182461&merchant_id=1407229&settlement_currency=&payment_id=176078290&product_id=&currency=USD&card_bin=537541&response_code=&card_type=MasterCard&amount=100&sender_email=vlad%40lesnoy.name&signature=4507069a57b0eb2e0aaa8e5d83d519c27efe5125")
	if parsing_err != nil {
		t.Fatal("Cannot parse URL query")
	}

	fmt.Println(parsedValues)
	finalResponse := parseRequestFormData(parsedValues)
	fmt.Printf("finalResponse %#v\n", finalResponse)
	fmt.Println(finalResponse.OrderTime)
}
