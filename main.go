package main

import (
	"go.ufukty.com/golits/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.New())
}
