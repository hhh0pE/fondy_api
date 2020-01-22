package fondy_api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/hhh0pE/easyjson"
	"github.com/pkg/errors"
)

func parseRequestFormData(values url.Values) FinalResponse {
	var fResp FinalResponse

	respReflect := reflect.ValueOf(&fResp)
	for i := 0; i < respReflect.Elem().NumField(); i++ {
		//field := respReflect.Type().Field(i)
		jsonTag := respReflect.Elem().Type().Field(i).Tag.Get("json")
		field := respReflect.Elem().Field(i)
		if val := values.Get(jsonTag); val != "" {

			if !field.CanAddr() {
				continue
			}
			typeName := field.Type().Name()

			if fieldUnmarshaller, ok := field.Addr().Interface().(json.Unmarshaler); ok {
				fieldUnmarshaller.UnmarshalJSON([]byte(val))
			} else {
				switch {
				case typeName == "string":
					field.SetString(val)
				case typeName == "int64":
					valInt64, _ := strconv.ParseInt(val, 10, 64)
					field.SetInt(valInt64)
				}
			}

		}
	}

	return fResp
}

func ReadFinalResponseFromRequest(r *http.Request) (*FinalResponse, error) {
	contentType := r.Header.Get("Content-Type")
	var finalResp FinalResponse
	if strings.HasSuffix(contentType, "x-www-form-urlencoded") {
		if parsing_form_err := r.ParseForm(); parsing_form_err != nil {
			return nil, errors.Wrap(parsing_form_err, "Error parsing request's form")
		}
		finalResp = parseRequestFormData(r.Form)
		return &finalResp, nil
	}
	return ReadFinalResponseFromReader(r.Body)
}

func ReadFinalResponseFromReader(reader io.Reader) (*FinalResponse, error) {
	inputBytes, read_err := ioutil.ReadAll(reader)
	if read_err != nil {
		return nil, errors.Wrap(read_err, "Cannot read data from reader")
	}

	responseBytes, _, _, getting_err := jsonparser.Get(inputBytes, "response")
	if getting_err != nil {
		return nil, errors.Wrap(getting_err, "Cannot get response from inputBytes")
	}
	if len(responseBytes) > 0 {
		inputBytes = responseBytes
	}

	var finalResp FinalResponse
	if err := easyjson.Unmarshal(inputBytes, &finalResp); err != nil {
		return nil, errors.Wrap(err, "Cannot unmarshal response")
	}

	return &finalResp, nil

}
