package brotli

import (
	"bytes"
	"errors"
	"io"

	"github.com/andybalholm/brotli"
	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
)

type brotliBlockCompressor struct{}

func (c *brotliBlockCompressor) CompressBlock(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := brotli.NewWriter(buf)
	n, err := w.Write(data)
	if n < len(data) {
		return nil, errors.New("short write")
	}
	if err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *brotliBlockCompressor) DecompressBlock(data []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(data))
	result, err := io.ReadAll(r)
	return result, err
}

func init() {
	goparquet.RegisterBlockCompressor(parquet.CompressionCodec_BROTLI, &brotliBlockCompressor{})
}
