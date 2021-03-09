package pkg

import (
	"context"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const LoggerRequestID = "RequestID"

const (
	ErrInternal                     = 0 // any internal error
	ErrDecode                       = 2 // failed to unmarshal incoming request
	ErrNoAccess                     = 3 // rejected by auth
	ErrExternalServerDoesNotRespond = 4 // external server does not respond
	ErrRequestValidation            = 5 // api server returns 400 error
	ErrInvalidCreditCard            = 6 // invalid credit card
	ErrRequiresCreditCardAction     = 7 // invalid credit card
)

type JSON map[string]interface{}

type key string

func createKey(name string) key {
	return key("OLARM_CTX_" + name)
}

var (
	RequestID        = createKey("REQUEST_ID")
	LoggerKey        = createKey("LOGGER_KEY")
	StripeCustomerID = createKey("STRIPE_CUSTOMER_ID")
	UserID           = createKey("USER_ID")
)

type uuidHex string

func EncodeHexUUID(id *uuid.UUID) uuidHex {
	b, _ := id.MarshalBinary()
	return uuidHex(hex.EncodeToString(b))
}

//GenerateNewStringUUID generates a uuid
func GenerateNewStringUUID() string {
	pid := uuid.New()
	return string(EncodeHexUUID(&pid))
}

//GetLogRequest returns a request logger
func GetLogRequest(r *http.Request) *logrus.Entry {
	log := r.Context().Value(LoggerKey)
	if log == nil {
		return logrus.WithField(LoggerRequestID, "not-set")
	}
	return log.(*logrus.Entry)
}

//GetLogContext returns a context logger
func GetLogContext(ctx context.Context) *logrus.Entry {
	log := ctx.Value(LoggerKey)
	if log == nil {
		return logrus.WithField(LoggerRequestID, "not-set")
	}
	return log.(*logrus.Entry)
}

// Error defines Error structure
type Error struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Extra interface{} `json:"extra"`

	OriginalError error `json:"-"`
}

// Error returns a error msg
func (e *Error) Error() string {
	return e.Msg
}
