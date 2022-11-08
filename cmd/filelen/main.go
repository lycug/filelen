// This must be package main
/*
	作为私有插件使用时，只用这种方式实现
	如果摒弃私有插件方式，可以只保留annalyzer.go文件的Newrun函数即可
*/
package main

import (
	"github.com/lycug/filelen/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
