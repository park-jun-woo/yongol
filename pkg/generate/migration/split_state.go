//ff:type feature=migration type=model
//ff:what splitState — splitStatements 의 single-char 상태 머신 (따옴표/블록코멘트/괄호 depth)
package migration

import "strings"

// splitState is the tiny state machine used by splitStatements to walk
// a SQL text while tracking string / comment context.
type splitState struct {
	out  []string
	sb   strings.Builder
	inSQ bool
	inDQ bool
	inBC bool
	depth int
}
