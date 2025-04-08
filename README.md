# golits (go literals)

golits is a CLI utility that lists all duplicate string literals found in a Go file. It is designed to catch multiple uses of same string literals for declaring different error types. Which confuse clients.

golits expects the first and only argument to be the filename. It exits with non-0 status code for IO and parsing errors and if there are multiple use of same string literal.

## Install

```sh
go install github.com/ufukty/golits
```

## Usage

```sh
$ golits errors.go
# "generic" (internal/mistakes/mistakes.go:15:27, internal/mistakes/mistakes.go:16:27, internal/mistakes/mistakes.go:17:27)
```

## Contribution

Issues, PRs and discussions are welcome.

## License

MIT
