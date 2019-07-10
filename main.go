package main

import (
	"github.com/wayneashleyberry/dontpanic/pkg/paniccheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(paniccheck.Analyzer)
}
