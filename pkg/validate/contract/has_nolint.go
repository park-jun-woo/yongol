//ff:func feature=validate-contract type=util control=iteration dimension=2 topic=preserve-safety
//ff:what hasNolint — 특정 라인(또는 바로 위 라인) 에 `// nolint:<rule>` 주석 존재 여부

package contract

import (
	"go/ast"
	"go/token"
	"strings"
)

// hasNolint reports whether file carries a `// nolint:<rule>` directive
// on line (or the line immediately above line). rule is matched
// case-insensitively — so the PRV-10 rule can ship as `// nolint:panic`
// (plan-mandated alias) while the generic form `// nolint:prv-13` works
// for every other rule.
//
// Multiple rules may be listed, separated by commas: `// nolint:prv-12,prv-17`.
func hasNolint(fset *token.FileSet, file *ast.File, line int, rule string) bool {
	if file == nil || fset == nil || line <= 0 {
		return false
	}
	rule = strings.ToLower(strings.TrimSpace(rule))
	for _, group := range file.Comments {
		for _, c := range group.List {
			pos := fset.Position(c.Slash)
			if pos.Line != line && pos.Line != line-1 {
				continue
			}
			if nolintMatches(c.Text, rule) {
				return true
			}
		}
	}
	return false
}
