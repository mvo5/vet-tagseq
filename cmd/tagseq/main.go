package main

import (
	"github.com/mvo5/vet-tagseq/tagseq"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(tagseq.Analyzer)
}
