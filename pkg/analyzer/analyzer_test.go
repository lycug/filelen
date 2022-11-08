package analyzer

// Tests for linters.

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, Analyzer, "test")
}

type linterSuite struct {
	suite.Suite
}

func (suite *linterSuite) TestContextLinter() {
	analysistest.Run(
		suite.T(), TestdataDir(),
		Analyzer, "/test/")
}

func TestLinterSuite(t *testing.T) {
	suite.Run(t, new(linterSuite))
}

func TestdataDir() string {
	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current test filename")
	}
	return filepath.Join(filepath.Dir(testFilename), "../testdata/src")
}
