// Command unsignedlint checks for potentially unsafe unsigned integer subtractions.
package main

import (
	"github.com/fingon/unsignedlint/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
