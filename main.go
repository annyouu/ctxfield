package main

import (
	"github.com/annyouu/ctxfield"
	// "golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	// singlechecker.Main(ctxfield.Analyzer)
	multichecker.Main(ctxfield.Analyzer)
}
