package system

import (
	"github.com/creasty/apperrors"
)

// ErrorCode represents error identifier specified for this application
type ErrorCode int

// Enum values for ErrorCode
const (
	ErrorUnknown ErrorCode = iota
	ErrorInvalidArgument
	ErrorUnauthorized
	ErrorForbbiden
	ErrorNotFound
	ErrorFailedToReadDB
	ErrorFailedToWriteDB
)

var stringByErrorCode = map[ErrorCode]string{
	ErrorUnknown:         "Unknown",
	ErrorInvalidArgument: "InvalidArgument",
	ErrorUnauthorized:    "Unauthorized",
	ErrorForbbiden:       "Forbbiden",
	ErrorNotFound:        "NotFound",
	ErrorFailedToReadDB:  "FailedToReadDB",
	ErrorFailedToWriteDB: "FailedToWriteDB",
}

// New returns an error with a status code
func (c ErrorCode) New(msg string) error {
	return apperrors.WithStatusCode(apperrors.New(msg), int(c))
}

// Wrap sets an error code to an error
func (c ErrorCode) Wrap(err error) error {
	return apperrors.WithStatusCode(apperrors.Wrap(err), int(c))
}

// WithReport annotates an error reportability
func (c ErrorCode) WithReport(err error) error {
	return apperrors.WithReport(c.Wrap(err))
}

func (c ErrorCode) String() string {
	return stringByErrorCode[c]
}
