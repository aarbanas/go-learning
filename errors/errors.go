package errors

import (
	"errors"
	"fmt"
)

func NegativeNumError() error {
	err := errors.New("math: can't calculate square root of negative number")

	return err
}

func ScanErrorMessage(err error) error {
	wrappedErr := fmt.Errorf("error on input: %w", err)

	return wrappedErr
}
