//ff:func feature=rule type=test-helper control=sequence
//ff:what newGround — populate_* 단위 테스트용 빈 maps 가 초기화된 *rule.Ground 생성

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// newGround returns a fully initialized *rule.Ground with empty maps so
// populate_* functions can write without nil-panics.
func newGround() *rule.Ground {
	return &rule.Ground{
		Lookup:  make(map[string]rule.StringSet),
		Types:   make(map[string]string),
		Pairs:   make(map[string]rule.StringSet),
		Config:  make(map[string]bool),
		Vars:    make(rule.StringSet),
		Flags:   make(rule.StringSet),
		Schemas: make(map[string][]string),
	}
}
