# Golits

<img src="assets/Logo@3x.png" width="256px">

Golits is a go-vet style analyzer.

Golits lists all the string literals that is mentioned more than once in a file.

Golits is designed to catch multiple uses of same string literals for declaring different error types. Which confuse clients.

## Install

```sh
go install github.com/ufukty/golits@latest
```

## Usage

Direct use:

```sh
$ golits ./...
```

Via Go Vet tool:

```sh
go vet --vettool="$(which golits)" ./...
```

## Example

```go
// file.go
package testdata

import "fmt"

var (
  ErrA = fmt.Errorf("a")
  ErrB = fmt.Errorf("a")
  ErrC = fmt.Errorf("c")
  ErrD = fmt.Errorf("d")
)

var (
  ErrE = fmt.Errorf("e")
  ErrF = fmt.Errorf("a")
  ErrG = fmt.Errorf("g")
  ErrH = fmt.Errorf("d")
)
```

```sh
$ golits .
file.go:6:20: duplicated string literal "a"
file.go:7:20: duplicated string literal "a"
file.go:14:20: duplicated string literal "a"
file.go:9:20: duplicated string literal "d"
file.go:16:20: duplicated string literal "d"
```

## License

MIT
