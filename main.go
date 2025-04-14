package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func lines(fs *token.FileSet, positions []token.Pos) string {
	ss := []string{}
	for _, pos := range positions {
		ss = append(ss, fs.Position(pos).String())
	}
	return strings.Join(ss, ", ")
}

func Main() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("expected filename as the first and only argument")
	}
	filename := os.Args[1]
	fh, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening %q: %w", filename, err)
	}
	defer fh.Close()

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, fh, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	order := []string{}
	occurrences := map[string][]token.Pos{}
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if bl, ok := n.(*ast.BasicLit); ok && bl.Kind == token.STRING {
			if !has(occurrences, bl.Value) {
				occurrences[bl.Value] = []token.Pos{}
				order = append(order, bl.Value)
			}
			occurrences[bl.Value] = append(occurrences[bl.Value], bl.Pos())
			return false
		}
		return true
	})

	duplicates := false
	for _, lit := range order {
		if os := occurrences[lit]; len(os) > 1 {
			fmt.Printf("%s (%s)\n", lit, lines(fs, os))
			duplicates = true
		}
	}

	if duplicates {
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
