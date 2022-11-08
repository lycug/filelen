package analyzer

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

var MaxLines int = 799

var Analyzer = &analysis.Analyzer{
	Name: "fll",
	Doc:  "max file line length",
	Run:  run,
}

func init() {
	Analyzer.Flags.IntVar(&MaxLines, "max-lines", 800, "max lines of file")
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		fs := pass.Fset.File(file.Pos())
		lines := fs.LineCount()
		if lines > MaxLines {
			pass.Report(analysis.Diagnostic{
				Pos:            file.Pos(),
				End:            0,
				Category:       "file lines length out of limit",
				Message:        fmt.Sprintf("file has %d lines which is out of limit %d lines", lines, MaxLines),
				SuggestedFixes: nil,
			})
		}
	}
	return nil, nil
}

func Newrun(pass *analysis.Pass, maxLineNum int) (interface{}, error) {
	if maxLineNum == 0 {
		maxLineNum = MaxLines
	}

	for _, file := range pass.Files {
		fs := pass.Fset.File(file.Pos())
		lines := fs.LineCount()
		if lines > maxLineNum {
			pass.Report(analysis.Diagnostic{
				Pos:            file.Pos(),
				End:            0,
				Category:       "file lines length out of limit",
				Message:        fmt.Sprintf("file has %d lines which is out of limit %d lines", lines, maxLineNum),
				SuggestedFixes: nil,
			})
		}
	}
	return nil, nil
}
