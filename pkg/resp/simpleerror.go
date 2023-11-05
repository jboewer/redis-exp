package resp

// SimpleError is a RESP simple error type.
type SimpleError struct {
	data  []byte
	Value string
}

func parseSimpleError(data []byte) (SimpleError, error) {
	return SimpleError{}, nil
}

func (e SimpleError) Type() TypeIdentifier {
	return TypeIdentifierSimpleError
}

func (e SimpleError) Bytes() []byte {
	return e.data
}
