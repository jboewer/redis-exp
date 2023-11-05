package resp

type TypeIdentifier byte

const (
	TypeIdentifierSimpleError TypeIdentifier = '-'
	TypeIdentifierInteger     TypeIdentifier = ':'
	TypeIdentifierBulkString  TypeIdentifier = '$'
	TypeIdentifierArray       TypeIdentifier = '*'

	crlf = "\r\n"
)

type Value interface {
	Type() TypeIdentifier
	Bytes() []byte
}
