package inspect

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// returns a map of literals and list of their occurrences
func occurrences(f *ast.File) map[string][]analysis.Range {
	occs := map[string][]analysis.Range{}
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if bl, ok := n.(*ast.BasicLit); ok && bl.Kind == token.STRING {
			if !has(occs, bl.Value) {
				occs[bl.Value] = []analysis.Range{}
			}
			occs[bl.Value] = append(occs[bl.Value], bl)
			return false
		}
		return true
	})
	return occs
}

// existence of duplicate entries is not counted as error
func file(pass *analysis.Pass, f *ast.File) {
	for lit, occs := range occurrences(f) {
		if len(occs) > 1 {
			for _, occ := range occs {
				pass.ReportRangef(occ, "duplicated string literal %s", lit)
			}
		}
	}
}

func Files(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		file(pass, f)
	}
	return nil, nil
}
