//ff:func feature=gen-gogin type=util control=sequence
//ff:what bodyReport — 단일 루프 body의 순수 라인 수 계산 후 PureLinesReport 반환

package qcheck

import (
	"go/ast"
	"go/token"
	"strings"
)

// bodyReport builds a PureLinesReport for one loop body slice. Lines between
// Lbrace and Rbrace (exclusive) are counted, skipping blank lines and lines
// starting with //. Multi-line block comments are not detected (an acceptable
// approximation for generator output which rarely carries them).
func bodyReport(fset *token.FileSet, fnName, kind string, headerPos token.Pos, body *ast.BlockStmt, src string) PureLinesReport {
	start := fset.Position(body.Lbrace).Line
	end := fset.Position(body.Rbrace).Line
	lines := strings.Split(src, "\n")
	pure := countPureLines(lines, start, end)
	return PureLinesReport{
		Func:      fnName,
		LoopKind:  kind,
		Line:      fset.Position(headerPos).Line,
		PureLines: pure,
	}
}
