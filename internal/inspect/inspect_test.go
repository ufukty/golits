package inspect

import (
	"fmt"
	"os"
)

func basicPrinter(duplicates []Entry) {
	for _, entry := range duplicates {
		fmt.Println(entry.LiteralValue)
		for _, occ := range entry.Occurrences {
			fmt.Println(" ", occ.String())
		}
	}
}

func ExampleFile_positive() {
	fh, _ := os.Open("testdata/positive.go")
	defer fh.Close()

	duplicates, _ := File(fh, "positive.go")
	basicPrinter(duplicates)

	// Output:
	// "a"
	//   positive.go:6:20
	//   positive.go:7:20
	//   positive.go:14:20
	// "d"
	//   positive.go:9:20
	//   positive.go:16:20
}
