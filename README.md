# parquet-go-block-compressors

<p align="center">
<a href="https://godoc.org/github.com/akrennmair/parquet-go-block-compressors"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?color=blue"></a>
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
)
```

## License

See the file `LICENSE` for further license information.

## Author

Andreas Krennmair <ak@synflood.at>