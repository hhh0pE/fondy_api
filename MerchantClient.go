package fondy_api

import (
	"net/http"
)

type Client struct {
	MerchantID int
	Password   string

	lastError error
}

func (c Client) LastError() error {
	return c.lastError
}

func (c *Client) GenerateHTTPRequest(r Request) *http.Request {
	req, err := generateHTTPRequest(r, c.MerchantID, c.Password)
	if err != nil {
		c.lastError = err
	}
	return req
}

func (c *Client) CreatePayment(r *CreatePaymentRequest) *URLResultResponse {
	res, err := SendRequest(c, r)
	if err != nil {
		c.lastError = err
	}
	if casted, ok := res.(*URLResultResponse); ok {
		return casted
	}
	return nil
}

func (c *Client) GetPaymentStatus(r *GetPaymentStatusRequest) *FinalResponse {
	res, err := SendRequest(c, r)
	if err != nil {
		c.lastError = err
	}
	if casted, ok := res.(*FinalResponse); ok {
		return casted
	}
	return nil
}

func (c *Client) CreateRecurringPayment(r *PaymentByToken) *FinalResponse {
	res, err := SendRequest(c, r)
	if err != nil {
		c.lastError = err
	}
	if casted, ok := res.(*FinalResponse); ok {
		return casted
	}
	return nil
}

func (c *Client) GetOrderStatus(orderID string) *FinalResponse {
	var ps GetPaymentStatusRequest
	ps.OrderID = orderID
	return c.GetPaymentStatus(&ps)
}
