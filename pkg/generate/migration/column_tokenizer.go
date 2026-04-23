//ff:type feature=migration type=model
//ff:what columnTokenizer — tokenizeColumnDef 의 단문자 상태 머신
package migration

import "strings"

// columnTokenizer is the state machine backing tokenizeColumnDef.
type columnTokenizer struct {
	out   []string
	sb    strings.Builder
	depth int
	inSQ  bool
	inDQ  bool
}
