package analyzer_test

import (
	"testing"

	"github.com/fingon/unsignedlint/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUnsignedLint(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "a") // "a" is the package name in testdata/src/a
}
