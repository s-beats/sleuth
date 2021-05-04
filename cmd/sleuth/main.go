package main

import (
	"github.com/sivchari/sleuth"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(sleuth.Analyzer) }
