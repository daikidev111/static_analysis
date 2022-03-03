package main

import (
	error_handling "github.com/daikidev111/static_analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(error_handling.Analyzer) }
