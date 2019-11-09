package fondy_api

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

func SignRequestBytes(merchantID int, password string, msgBytes *[]byte) error {
	if msgBytes == nil {
		return errors.New("nil msgBytes")
	}

	*msgBytes, _ = jsonparser.Set(*msgBytes, []byte(strconv.Itoa(merchantID)), "merchant_id")

	var urlData = make(url.Values)
	iterating_err := jsonparser.ObjectEach(*msgBytes, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		urlData.Add(string(key), string(value))
		return nil
	})

	if iterating_err != nil {
		return errors.Wrap(iterating_err, "GenerateSignature jsonparser.ObjectEach error")
	}

	var keys = make([]string, len(urlData))
	var ki = 0
	for key, _ := range urlData {
		keys[ki] = key
		ki++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var signValues []string
	signValues = append(signValues, password)
	for _, k := range keys {
		signValues = append(signValues, urlData.Get(k))
	}

	strToSign := strings.Join(signValues, "|")

	sig := sha1.New()
	fmt.Fprint(sig, strToSign)
	signature := hex.EncodeToString(sig.Sum(nil))

	*msgBytes, _ = jsonparser.Set(*msgBytes, []byte(`"`+signature+`"`), "signature")
	return nil
}

//func GenerateSignature(password string, requestParams url.Values) string {
//	fmt.Println("GenerateSignature input", requestParams)
//	var keys = make([]string, len(requestParams))
//	var ki = 0
//	for key, _ := range requestParams {
//		keys[ki] = key
//		ki++
//	}
//	sort.Slice(keys, func(i, j int) bool {
//		return keys[i] < keys[j]
//	})
//
//	var signValues []string
//	signValues = append(signValues, password)
//	for _, k := range keys {
//		signValues = append(signValues, requestParams.Get(k))
//	}
//
//	strToSign := strings.Join(signValues, "|")
//	fmt.Println("strToSign", strToSign)
//	sig := sha1.New()
//	fmt.Fprint(sig, strToSign)
//	return hex.EncodeToString(sig.Sum(nil))
//}
