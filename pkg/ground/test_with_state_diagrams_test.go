//ff:func feature=rule type=test-helper control=sequence
//ff:what withStateDiagrams — Mermaid StateDiagram 슬라이스를 Fullstack.StateDiagrams 에 append 하는 option

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withStateDiagrams attaches Mermaid state diagrams.
func withStateDiagrams(sds ...*statemachine.StateDiagram) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.StateDiagrams = append(fs.StateDiagrams, sds...) }
}
