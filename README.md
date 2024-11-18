# rdb-go

Parse Redis RDB dump files. This library is based on [redis-rdb-tools](https://github.com/sripathikrishnan/redis-rdb-tools).

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tommy351/rdb-go)](https://github.com/Vergangenheit/rdb-go/releases) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/Vergangenheit/rdb-go) ![Test](https://github.com/Vergangenheit/rdb-go/workflows/Test/badge.svg) [![codecov](https://codecov.io/gh/tommy351/rdb-go/branch/master/graph/badge.svg)](https://codecov.io/gh/tommy351/rdb-go)

## Install

This library can be used as a package.

```sh
go get github.com/Vergangenheit/rdb-go
```

Or as a command line tool.

```sh
go get github.com/Vergangenheit/rdb-go/cmd/rdb
rdb path/to/dump.rdb
```

## Usage

Use `Parser` to iterate over a RDB dump file.

```go
import (
  rdb "github.com/Vergangenheit/rdb-go"
)

parser := NewParser(file)

for {
  data, err := parser.Next()

  if err == io.EOF {
    break
  }

  if err != nil {
    panic(err)
  }

  // ...
}
```

See examples in the [documentation](https://pkg.go.dev/github.com/Vergangenheit/rdb-go) or [cmd/rdb/main.go](cmd/rdb/main.go) for more details.

## License

MIT
