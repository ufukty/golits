package inspect

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

type Entry struct {
	LiteralValue string
	Occurrences  []token.Position
}

// existence of duplicate entries is not counted as error
func File(fh io.Reader, filename string) ([]Entry, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, fh, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parsing: %w", err)
	}

	order := []string{}
	occurrences := map[string][]token.Position{}
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if bl, ok := n.(*ast.BasicLit); ok && bl.Kind == token.STRING {
			if !has(occurrences, bl.Value) {
				occurrences[bl.Value] = []token.Position{}
				order = append(order, bl.Value)
			}
			occurrences[bl.Value] = append(occurrences[bl.Value], fs.Position(bl.Pos()))
			return false
		}
		return true
	})

	duplicates := []Entry{}
	for _, lit := range order {
		if len(occurrences[lit]) > 1 {
			duplicates = append(duplicates, Entry{
				LiteralValue: lit,
				Occurrences:  occurrences[lit],
			})
		}
	}

	return duplicates, nil
}
