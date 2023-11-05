package resp

// region Integer

// Integer is a RESP integer type.
type Integer struct {
	data  []byte
	Value int
}

func parseInteger(data []byte) (Integer, error) {
	return Integer{}, nil
}

func (i Integer) Type() TypeIdentifier {
	return TypeIdentifierInteger
}

func (i Integer) Bytes() []byte {
	return i.data
}

//endregion
