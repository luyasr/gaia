package errors

import (
	goerrors "errors"
)

func Unwrap(err error) error {
	return goerrors.Unwrap(err)
}

func Is(err, target error) bool {
	return goerrors.Is(err, target)
}

func As(err error, target any) bool {
	return goerrors.As(err, &target)
}
