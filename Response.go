package fondy_api

import (
	"github.com/hhh0pE/easyjson"
)

type Response interface {
	easyjson.Unmarshaler
	ResponseError() error
	UnwrappedObject() interface{}
}
