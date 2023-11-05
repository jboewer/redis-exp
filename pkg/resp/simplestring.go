package resp

import (
	"fmt"
	"strings"
)

const TypeIdentifierSimpleString TypeIdentifier = '+'

// SimpleString is a RESP simple string type.
type SimpleString struct {
	data  []byte
	Value string
}

func NewSimpleString(stringVal string) (SimpleString, error) {
	var s SimpleString

	if strings.Contains(stringVal, crlf) {
		return s, fmt.Errorf("simple string cannot contain CRLF")
	}

	s.data = []byte(fmt.Sprintf("+%s%s", stringVal, crlf))
	s.Value = stringVal

	return s, nil
}

func (s SimpleString) Type() TypeIdentifier {
	return TypeIdentifierSimpleString
}

func (s SimpleString) Bytes() []byte {
	return s.data
}
