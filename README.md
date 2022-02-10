# parquet-go-block-compressors

<p align="center">
<a href="https://github.com/akrennmair/parquet-go-block-compressors/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache%202-blue"></a>
</p>

This library implements more compression algorithms for [github.com/fraugster/parquet-go](github.com/fraugster/parquet-go). By default,
`parquet-go` library only supports `GZIP` and `SNAPPY` as compression algorithms to minimize the list
of dependencies.

All you need to do is import the correct package into your program and the respective compression
algorithm will be automatically available in `parquet-go`.

```go
import (
    _ "github.com/akrennmair/parquet-go-block-compressors/zstd" // registers the Zstd block compressor with parquet-go
    _ "github.com/akrennmair/parquet-go-block-compressors/brotli" // registers the Brotli block compressor with parquet-go
    _ "github.com/akrennmair/parquet-go-block-compressors/lzo" // registers the LZO block compressor with parquet-go
    _ "github.com/akrennmair/parquet-go-block-compressors/lz4raw" // registers the LZ4 block compressor with the LZ4_RAW compression type with parquet-go
)
```

## Documentation

* [Go documentation for github.com/akrennmair/parquet-go-block-compressors/zstd](https://pkg.go.dev/github.com/akrennmair/parquet-go-block-compressors/zstd)
* [Go documentation for github.com/akrennmair/parquet-go-block-compressors/brotli](https://pkg.go.dev/github.com/akrennmair/parquet-go-block-compressors/brotli)
* [Go documentation for github.com/akrennmair/parquet-go-block-compressors/lzo](https://pkg.go.dev/github.com/akrennmair/parquet-go-block-compressors/lzo)
* [Go documentation for github.com/akrennmair/parquet-go-block-compressors/lz4raw](https://pkg.go.dev/github.com/akrennmair/parquet-go-block-compressors/lz4raw)

## License

See the file `LICENSE` for further license information for the libraries contained in this repository.

Please note that `github.com/akrennmair/parquet-go-block-compressors/lzo` is built using `github.com/cyberdelia/lzo` which in turn
uses the original [LZO implementation which is licensed as GPLv2+](http://www.oberhumer.com/opensource/lzo/). Please be aware of the licensing implication this can have if you intend to use this in closed-source products that you intend to distribute.

## Author

Andreas Krennmair <ak@synflood.at>
