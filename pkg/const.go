package pkg

import (
	"context"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	ErrInternal                     = 0 // any internal error
	ErrDecode                       = 1 // failed to unmarshal incoming request
	ErrExternalServerDoesNotRespond = 2 // external server does not respond
	ErrRequestValidation            = 3 // api server returns 400 error
)

type JSON map[string]interface{}

var (
	RequestID = "REQUEST_ID"
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
	log := r.Context().Value(RequestID)
	if log == nil {
		return logrus.WithField(RequestID, "not-set")
	}
	return log.(*logrus.Entry)
}

//GetLogContext returns a context logger
func GetLogContext(ctx context.Context) *logrus.Entry {
	log := ctx.Value(RequestID)
	if log == nil {
		return logrus.WithField(RequestID, "not-set")
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
