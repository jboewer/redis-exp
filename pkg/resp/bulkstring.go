package resp

import "fmt"

// region BulkString

// BulkString is a RESP bulk string type.
type BulkString struct {
	data   []byte
	Length int
	Value  string
}

func parseBulkString(data []byte) (BulkString, error) {
	// example: $6\r\nfoobar\r\n
	var bs BulkString

	if data[0] != byte(TypeIdentifierBulkString) {
		return bs, fmt.Errorf("invalid bulk string type identifier: %v", data[0])
	}

	_, err := fmt.Sscanf(string(data[1:]), "%d\r\n", &bs.Length)
	if err != nil {
		return bs, err
	}

	valueOffset := len(fmt.Sprintf("%d", bs.Length)) + 3
	bs.Value = string(
		data[valueOffset : bs.Length+valueOffset],
	)

	bs.data = data[:bs.Length+valueOffset+len(crlf)]

	return bs, nil
}

func NewBulkString(s string) BulkString {
	data := []byte(fmt.Sprintf("$%d%s%s%s", len(s), crlf, s, crlf))

	return BulkString{
		data:   data,
		Length: len(s),
		Value:  s,
	}
}

func (b BulkString) Type() TypeIdentifier {
	return TypeIdentifierBulkString
}

func (b BulkString) Bytes() []byte {
	return b.data
}

//endregion
