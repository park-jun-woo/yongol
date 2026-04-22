//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what parseComments — extracts v2 sequences from a comment list
package ssac

import (
	"go/ast"
	"go/token"
	"strings"
)

// parseComments extracts v2 sequences from a list of comments.
// When fset is non-nil, each sequence is annotated with its 1-based comment line number.
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
