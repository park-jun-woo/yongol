//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what hasStateNeutralComment — detects the `// @state-neutral` function-level annotation

package ssac

import (
	"go/ast"
	"strings"
)

// hasStateNeutralComment reports whether any line in the given comment slice
// matches the bare "@state-neutral" directive. The annotation declares that
// the function is intentionally independent of the target resource's state
// machine; XSM-27 skips functions marked this way.
func hasStateNeutralComment(comments []*ast.Comment) bool {
	for _, c := range comments {
		line := strings.TrimPrefix(c.Text, "//")
		line = strings.TrimSpace(line)
		if line == "@state-neutral" {
			return true
		}
	}
	return false
}
