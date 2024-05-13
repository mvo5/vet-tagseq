package tagseq_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mvo5/vet-tagseq/tagseq"
)

func TestAnalyzerHappy(t *testing.T) {
	testdata := analysistest.TestData()
	// see anaysistest/ subdirs, the expected messages are embedded there
	analysistest.Run(t, testdata, tagseq.Analyzer, "happy", "sad")
}
