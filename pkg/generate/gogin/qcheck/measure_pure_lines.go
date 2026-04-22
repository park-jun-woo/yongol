//ff:func feature=gen-gogin type=util control=sequence
//ff:what MeasurePureLines — range/for 루프 body 의 순수 라인 수 계산 (filefunc Q4 근거)

package qcheck

import (
	"go/parser"
	"go/token"
)

// MeasurePureLines parses src and returns one PureLinesReport per loop found
// in any top-level function. Blank lines and pure-comment lines are excluded
// from the count (matching filefunc's Q4 semantics).
func MeasurePureLines(filename, src string) ([]PureLinesReport, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return collectFileLoopReports(fset, file, src), nil
}
