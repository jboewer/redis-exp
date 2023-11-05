package resp

import "fmt"

// Array is a RESP array type.
type Array struct {
	data   []byte
	Length int
	Value  []Value
}

func NewArray(values []Value) (Array, error) {
	var arr Array

	arr.Value = values
	arr.Length = len(values)

	data := fmt.Sprintf("*%d%s", arr.Length, crlf)

	for _, v := range values {
		data += string(v.Bytes())
	}

	arr.data = []byte(data)

	return arr, nil
}

func (a Array) Type() TypeIdentifier {
	return TypeIdentifierArray
}

func (a Array) Bytes() []byte {
	return a.data
}
