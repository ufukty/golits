package testdata

import "fmt"

var (
	ErrA = fmt.Errorf("a") // want `duplicated string literal`
	ErrB = fmt.Errorf("a") // want `duplicated string literal`
	ErrC = fmt.Errorf("c")
	ErrD = fmt.Errorf("d") // want `duplicated string literal`
)

var (
	ErrE = fmt.Errorf("e")
	ErrF = fmt.Errorf("a") // want `duplicated string literal`
	ErrG = fmt.Errorf("g")
	ErrH = fmt.Errorf("d") // want `duplicated string literal`
)
