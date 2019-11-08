package fondy_api

import (
	"net/http"

	"bytes"

	"io"

	"github.com/hhh0pE/easyjson"
	"github.com/pkg/errors"
)

type Request interface {
	PrepareData() error
	URLSuffix() string
	ResponseObject() Response
	easyjson.Marshaler
}

func generateHTTPRequest(r Request, merchantID int, password string) (*http.Request, error) {
	if r == nil {
		return nil, errors.New("CreatePaymentRequest is nil")
	}

	if password == "" {
		return nil, errors.New("MerchantPassword is required")
	}
	if merchantID == 0 {
		return nil, errors.New("MerchantID is required")
	}

	if err := r.PrepareData(); err != nil {
		return nil, errors.Wrap(err, "Error preparing Request")
	}

	marshalledBytes, marshall_err := easyjson.Marshal(r)
	if marshall_err != nil {
		return nil, errors.Wrap(marshall_err, "GenerateSignature marshalling error")
	}

	if err := SignRequestBytes(merchantID, password, &marshalledBytes); err != nil {
		return nil, errors.Wrap(err, "Error signing Request")
	}

	marshalledBytes = append([]byte(`{"request":`), append(marshalledBytes, []byte("}")...)...)

	req, err := http.NewRequest("POST", FondyURL+r.URLSuffix(), bytes.NewReader(marshalledBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil

}

func SendRequest(c *Client, r Request) (interface{}, error) {
	c.lastError = nil
	req, generate_err := generateHTTPRequest(r, c.MerchantID, c.Password)
	if generate_err != nil {
		return nil, generate_err
	}

	resp, req_err := http.DefaultClient.Do(req)
	if req_err != nil {
		return nil, req_err
	}

	var respBuff bytes.Buffer
	reader := io.TeeReader(resp.Body, &respBuff)
	respObject := r.ResponseObject()
	if err := easyjson.UnmarshalFromReader(reader, respObject); err != nil {
		return nil, errors.Wrap(err, "Cannot unmarshal response")
	}

	if respObject.ResponseError() != nil {
		return nil, respObject.ResponseError()
	}

	return respObject.UnwrappedObject(), nil
}
