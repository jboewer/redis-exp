package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSimpleString(t *testing.T) {
	type test struct {
		data string
		want SimpleString
	}

	tests := []test{
		{
			data: "OK",
			want: SimpleString{
				data:  []byte("+OK\r\n"),
				Value: "OK",
			},
		},
	}

	for _, tc := range tests {
		got, _ := NewSimpleString(tc.data)
		assert.Equal(t, tc.want.Value, got.Value)
	}
}

func TestParseBulkString(t *testing.T) {
	type test struct {
		data []byte
		want BulkString
	}

	tests := []test{
		{
			data: []byte("$6\r\nfoobar\r\n"),
			want: BulkString{
				data:   []byte("$6\r\nfoobar\r\n"),
				Length: 6,
				Value:  "foobar",
			},
		},
		{
			data: []byte("$6\r\nfoobar\r\nfoooooo"),
			want: BulkString{
				data:   []byte("$6\r\nfoobar\r\n"),
				Length: 6,
				Value:  "foobar",
			},
		},
	}

	for _, tc := range tests {
		got, _ := parseBulkString(tc.data)
		assert.Equal(t, tc.want.data, got.data)
		assert.Equal(t, tc.want.Length, got.Length)
		assert.Equal(t, tc.want.Value, got.Value)
	}
}

func TestNewBulkString(t *testing.T) {
	type test struct {
		data string
		want BulkString
	}

	tests := []test{
		{
			data: "foobar",
			want: BulkString{
				data:   []byte("$6\r\nfoobar\r\n"),
				Length: 6,
				Value:  "foobar",
			},
		},
	}

	for _, tc := range tests {
		got, _ := NewBulkString(tc.data)
		assert.Equal(t, tc.want.Length, got.Length)
		assert.Equal(t, tc.want.Value, got.Value)
		assert.Equal(t, tc.want.data, got.data)
	}
}
