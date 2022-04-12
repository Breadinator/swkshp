package errors

import "errors"

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target *error) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func New[S ~string](text S) error {
	return errors.New(string(text))
}
