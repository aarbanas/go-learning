package errors

import "errors"

func NegativeNumError() error {
	err := errors.New("math: can't calculate square root of negative number")

	return err
}
