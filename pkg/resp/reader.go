package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Reader struct {
	rd  *bufio.Reader
	buf []byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		rd:  bufio.NewReader(r),
		buf: make([]byte, 4096),
	}
}

func (r *Reader) ReadValue() (Value, error) {
	typeByte, err := r.rd.ReadByte()
	if err != nil {
		return nil, err
	}

	valueType := TypeIdentifier(typeByte)

	switch valueType {
	case TypeIdentifierSimpleString:
		l, err := r.readLine()
		if err != nil {
			return nil, err
		}
		return NewSimpleString(string(l))
	case TypeIdentifierSimpleError:
		l, err := r.readLine()
		if err != nil {
			return nil, err
		}
		return NewSimpleError(string(l))
	case TypeIdentifierBulkString:
		length, err := r.readInt()
		if err != nil {
			return nil, err
		}
		data := make([]byte, length+len(crlf)) // Read including CRLF
		_, err = r.rd.Read(data)
		if err != nil {
			return nil, err
		}

		return NewBulkString(string(data[:length])), nil // Strip CRLF
	case TypeIdentifierArray:
		length, err := r.readInt()
		if err != nil {
			return nil, err
		}
		values := make([]Value, length)
		for i := 0; i < length; i++ {
			values[i], err = r.ReadValue()
			if err != nil {
				return nil, err
			}
		}
		return NewArray(values), nil
	default:
		return nil, fmt.Errorf("unknown type identifier: %v", valueType)
	}
}

func (r *Reader) readLine() ([]byte, error) {
	var fullLine []byte

	for {
		currentLine, b, err := r.rd.ReadLine()
		if err != nil {
			return nil, err
		}

		fullLine = append(fullLine, currentLine...)

		if !b {
			break
		}
	}

	return fullLine, nil
}

func (r *Reader) readInt() (int, error) {
	intLine, err := r.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(intLine))
}
