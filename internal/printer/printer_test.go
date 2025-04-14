package printer

import (
	"os"

	"github.com/ufukty/golits/internal/inspect"
)

func ExampleOneLine_positive() {
	fh, _ := os.Open("testdata/positive.go")
	defer fh.Close()

	duplicates, _ := inspect.File(fh, "positive.go")
	OneLine(duplicates)

	// Output:
	// "a" (positive.go:6:20, positive.go:7:20, positive.go:14:20)
	// "d" (positive.go:9:20, positive.go:16:20)
}
