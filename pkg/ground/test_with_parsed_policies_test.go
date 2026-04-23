//ff:func feature=rule type=test-helper control=sequence
//ff:what withParsedPolicies — Rego Policy 슬라이스를 Fullstack.ParsedPolicies 에 append 하는 option

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withParsedPolicies attaches Rego policies.
func withParsedPolicies(pols ...rego.Policy) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ParsedPolicies = append(fs.ParsedPolicies, pols...) }
}
