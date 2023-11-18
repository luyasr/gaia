package errors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	var err *Error
	err = New(401, "reason", "message")
	err2 := err.WithMetadata(map[string]string{"foo": "bar"}).WithMetadata(map[string]string{"bar": "baz"})
	err3 := err2.WithCause(errors.New("cause"))

	t.Log(err3)
}

func TestFromError(t *testing.T) {
	err := errors.New("my_error")

	fromError := FromError(err)
	t.Log(fromError)
}
