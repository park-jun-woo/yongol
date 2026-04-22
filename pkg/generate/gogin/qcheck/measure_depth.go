//ff:func feature=gen-gogin type=util control=sequence
//ff:what MeasureDepth — Go 소스의 함수별 최대 nesting depth 계산 (filefunc Q1 근거)

package qcheck

import (
	"go/parser"
	"go/token"
)

// MeasureDepth parses src as a Go file and returns one DepthReport per
// top-level function declaration. Parser errors propagate unchanged so
// callers can surface template bugs (malformed emit) as ERROR.
func MeasureDepth(filename, src string) ([]DepthReport, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	return collectDepthReports(file), nil
}
