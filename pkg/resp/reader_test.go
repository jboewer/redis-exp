package resp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader_ReadValueSimpleString(t *testing.T) {
	type test struct {
		data []byte
		want Value
	}

	tests := []test{
		{
			data: []byte("+OK\r\n"),
			want: SimpleString{
				data:  []byte("+OK\r\n"),
				Value: "OK",
			},
		},
	}

	for _, tc := range tests {
		byteReader := bytes.NewReader(tc.data)
		respReader := NewReader(byteReader)

		value, err := respReader.ReadValue()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		got, ok := value.(SimpleString)
		if !ok {
			t.Fatalf("Expected SimpleString, got %T", value)

		}

		assert.Equal(t, tc.want, got)
	}
}

func TestReader_ReadValue(t *testing.T) {
	type test struct {
		data []byte
		want Value
	}

	tests := []test{
		{
			data: []byte("+OK\r\n"),
			want: SimpleString{
				data:  []byte("+OK\r\n"),
				Value: "OK",
			},
		},
		{
			data: []byte("$6\r\nfoobar\r\n"),
			want: BulkString{
				data:   []byte("$6\r\nfoobar\r\n"),
				Length: 6,
				Value:  "foobar",
			},
		},
		{
			data: []byte("*0\r\n"),
			want: Array{
				data:   []byte("*0\r\n"),
				Length: 0,
				Values: []Value{},
			},
		},
		{
			data: []byte("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"),
			want: Array{
				data:   []byte("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"),
				Length: 2,
				Values: []Value{
					BulkString{
						data:   []byte("$3\r\nfoo\r\n"),
						Length: 3,
						Value:  "foo",
					},
					BulkString{
						data:   []byte("$3\r\nbar\r\n"),
						Length: 3,
						Value:  "bar",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		byteReader := bytes.NewReader(tc.data)
		respReader := NewReader(byteReader)

		value, err := respReader.ReadValue()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, tc.want, value)
	}
}
