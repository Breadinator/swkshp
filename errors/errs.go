package errors

import (
	"errors"
	"fmt"
)

var (
	ErrGameNotFound     = errors.New("game not found")
	ErrParsingFailed    = errors.New("error while parsing")
	ErrResourceNotReady = errors.New("resource not ready")
	ErrHTTP             = errors.New("http error")
	ErrIllegalPath      = errors.New("illegal file path")
	ErrGameUnavailable  = errors.New("game not avalable or server is down")
	ErrType             = errors.New("invalid error")
)

func Wrap(err error, info string, a ...any) error {
	return fmt.Errorf(`%w: %s`, err, fmt.Sprintf(info, a...))
}
