//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 주석 리스트에서 v2 시퀀스를 추출
package ssac

import (
	"go/ast"
	"go/token"
	"strings"
)

// parseComments는 주석 리스트에서 v2 시퀀스를 추출한다.
// fset이 nil이 아니면 각 시퀀스에 주석 줄 번호(1-based)를 채운다.
func parseComments(fset *token.FileSet, comments []*ast.Comment) ([]Sequence, error) {
	cp := &commentParser{}
	for _, c := range comments {
		line := strings.TrimPrefix(c.Text, "//")
		line = strings.TrimSpace(line)
		commentLine := 0
		if fset != nil {
			commentLine = fset.Position(c.Slash).Line
		}
		if err := cp.processLine(line, commentLine); err != nil {
			return nil, err
		}
	}
	return cp.sequences, nil
}
