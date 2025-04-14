package printer

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/ufukty/golits/internal/inspect"
)

func join(occurrences []token.Position) string {
	ss := []string{}
	for _, pos := range occurrences {
		ss = append(ss, pos.String())
	}
	return strings.Join(ss, ", ")
}

func OneLine(duplicates []inspect.Entry) {
	for _, lit := range duplicates {
		fmt.Printf("%s (%s)\n", lit.LiteralValue, join(lit.Occurrences))
	}
}
