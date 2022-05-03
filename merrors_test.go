package merrors_test

import (
	"fmt"
	"testing"

	"github.com/kevin2027/merrors"
)

func TestString(t *testing.T) {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%s\n", merrors.String(err))
		}
	}()
	err = Caller()
	if err != nil {
		err = fmt.Errorf("new error: %w", err)
	}
}

func Caller() error {
	err := fmt.Errorf("error call")
	return err
}
