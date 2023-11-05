package resp

import (
	"fmt"
	"strings"
)

// SimpleError is a RESP simple error type.
type SimpleError struct {
	data  []byte
	Value string
}

func NewSimpleError(stringVal string) (SimpleError, error) {
	var err SimpleError

	if strings.Contains(stringVal, crlf) {
		return err, fmt.Errorf("simple error cannot contain CRLF")
	}

	err.data = []byte(fmt.Sprintf("%c%s%s", TypeIdentifierSimpleError, stringVal, crlf))
	err.Value = stringVal

	return err, nil
}

func (e SimpleError) Type() TypeIdentifier {
	return TypeIdentifierSimpleError
}

func (e SimpleError) Bytes() []byte {
	return e.data
}
