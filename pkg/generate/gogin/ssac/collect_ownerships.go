//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectOwnerships — fs.ParsedPolicies 전체에서 @ownership 매핑을 평탄화

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectOwnerships flattens every `@ownership` annotation parsed from the
// project's Rego policy files into a single slice. buildAuth consumes this
// to locate the mapping that matches a given `@auth <action> <resource>`
// sequence so the corresponding OwnerLookup<Resource> sqlc call can be
// emitted alongside the authz.Check.
//
// Ordering follows ParsedPolicies iteration, then per-policy ownership
// order — deterministic because both slices are deterministic upstream.
// The slice is safe to share across methodGen instances; no entry is
// mutated after parse.
func collectOwnerships(fs *yongol.Fullstack) []rego.OwnershipMapping {
	if fs == nil {
		return nil
	}
	var out []rego.OwnershipMapping
	for _, p := range fs.ParsedPolicies {
		out = append(out, p.Ownerships...)
	}
	return out
}
