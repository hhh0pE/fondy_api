package fondy_api

import (
	"io"
	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/hhh0pE/easyjson"
	"github.com/pkg/errors"
)

func ReadFinalResponseFromReader(reader io.Reader) (*FinalResponse, error) {

	inputBytes, _ := ioutil.ReadAll(reader)

	responseBytes, _, _, _ := jsonparser.Get(inputBytes, "response")
	if len(responseBytes) > 0 {
		inputBytes = responseBytes
	}

	var finalResp FinalResponse
	if err := easyjson.Unmarshal(inputBytes, &finalResp); err != nil {
		return nil, errors.Wrap(err, "Cannot unmarshal response")
	}

	return &finalResp, nil

}
