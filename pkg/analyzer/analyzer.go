package analyzer

import (
	"go.ufukty.com/golits/internal/inspect"
	"golang.org/x/tools/go/analysis"
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "golits",
		Doc:  "Lists the string literals declared multiple times in a file.",
		Run:  inspect.Files,
	}
}
