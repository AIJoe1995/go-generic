package errs

import (
	"errors"
	"fmt"
)

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("index %d out of range for length %d  invalid index.\n", index, length)
}

func NewErrKeyNotExist() error {
	return errors.New("key does not exist.\n")
}

func NewSystemError() error {
	return errors.New("system error. \n")
}
