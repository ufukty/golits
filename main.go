package main

import (
	"fmt"
	"os"

	"github.com/ufukty/golits/internal/inspect"
	"github.com/ufukty/golits/internal/printer"
)

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

	duplicates, err := inspect.File(fh, filename)
	if err != nil {
		return fmt.Errorf("inspecting: %w", err)
	}

	if len(duplicates) > 0 {
		printer.OneLine(duplicates)
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
