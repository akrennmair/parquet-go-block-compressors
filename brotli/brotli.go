package brotli

import (
	"bytes"
	"errors"
	"io"

	"github.com/andybalholm/brotli"
)

type BrotliBlockCompressor struct{}

func NewBrotliBlockCompressor() *BrotliBlockCompressor {
	return &BrotliBlockCompressor{}
}

func (c *BrotliBlockCompressor) CompressBlock(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := brotli.NewWriter(buf)
	n, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, errors.New("short write")
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *BrotliBlockCompressor) DecompressBlock(data []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(data))
	result, err := io.ReadAll(r)
	return result, err
}
