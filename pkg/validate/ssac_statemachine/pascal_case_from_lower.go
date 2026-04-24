//ff:func feature=validate type=util control=sequence topic=states
//ff:what pascalCaseFromLower — 소문자 단일 단어 리소스명 → PascalCase

package ssac_statemachine

import (
	"strings"
)

// pascalCaseFromLower returns "workflow" → "Workflow" (single-word resources
// only; multi-word diagram IDs already carry their own Symbol on StateDiagram).
func pascalCaseFromLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
