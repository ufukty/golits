package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestInspect(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), New(), "")
}
