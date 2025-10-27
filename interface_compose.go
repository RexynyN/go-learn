package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello high value software engineer")
	hashAndBroadcast(NewHashReader(payload))
}

type hashReader interface {
	io.Reader
	hash() string
}

type HashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *HashReader {
	return &HashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *HashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r hashReader) error {
	hash := r.hash()

	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("String of the bytes: ", string(b))
	return nil
}
