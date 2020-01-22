package fondy_api

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type FondyInt int64

func (sn FondyInt) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(sn), 10)), nil
}

func (sn *FondyInt) UnmarshalJSON(val []byte) error {
	valStr := strings.Trim(string(val), `"`)

	if valStr == "" {
		*sn = 0
		return nil
	}

	parsedNum, parsing_err := strconv.ParseInt(valStr, 10, 64)
	if parsing_err != nil {
		return errors.Wrap(parsing_err, "FondyInt ParseInt error")
	}
	*sn = FondyInt(parsedNum)

	return nil
}

type FondyFloat float64

func (sn *FondyFloat) UnmarshalJSON(val []byte) error {
	valStr := strings.Trim(string(val), `"`)

	if valStr == "" {
		*sn = 0
		return nil
	}

	parsedNum, parsing_err := strconv.ParseInt(valStr, 10, 64)
	if parsing_err != nil {
		return errors.Wrap(parsing_err, "FondyFloat ParseInt error")
	}

	*sn = FondyFloat(convertFondyNumberToFloat(parsedNum))

	return nil
}
func (sn FondyFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.0f", float64(sn)*100)), nil
}

func convertFondyNumberToFloat(val int64) float64 {
	var res float64
	res = float64(val) / 100
	return res
}

type FondyBool bool

func (sn *FondyBool) UnmarshalJSON(val []byte) error {
	valStr := strings.Trim(string(val), `"`)

	switch valStr {
	case "Y":
		*sn = true
	case "N":
		*sn = false
	case "":
		*sn = false
	default:
		return errors.New("Error unmarshaling FondyBool (" + valStr + ")")
	}

	return nil
}
func (sn FondyBool) MarshalJSON() ([]byte, error) {
	if sn {
		return []byte(`"Y"`), nil
	}
	return []byte(`"N"`), nil
}

type FondySeconds time.Duration

func (sn *FondySeconds) UnmarshalJSON(val []byte) error {
	valStr := strings.Trim(string(val), `"`)

	parsedNum, parsing_err := strconv.ParseInt(valStr, 10, 64)
	if parsing_err != nil {
		return errors.Wrap(parsing_err, "FondySeconds ParseInt error")
	}
	*sn = FondySeconds(time.Second * time.Duration(parsedNum))

	return nil
}
func (sn FondySeconds) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(time.Duration(sn)/time.Second), 10)), nil
}

type FondyTime struct {
	time.Time
}

func (sn *FondyTime) UnmarshalJSON(val []byte) error {
	valStr := strings.Trim(string(val), `"`)

	if valStr == "" {
		return nil
	}
	parsedTime, parsing_err := time.Parse("02.01.2006 15:04:05", valStr)
	if parsing_err != nil {
		return errors.Wrap(parsing_err, "FondyTime time.Parse("+valStr+") error")
	}
	sn.Time = parsedTime

	return nil
}
func (sn FondyTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + sn.Format("02.01.2006 15:04:05") + `"`), nil
}
