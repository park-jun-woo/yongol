//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XQP-30 — @ownership 매핑은 대응 sqlc 쿼리 OwnerLookup<Resource> 가 존재해야 함

package query_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqp30OwnerLookupQuery validates XQP-30: every Rego `@ownership` annotation
// demands a sqlc query whose name follows the `OwnerLookup<Resource>`
// convention declared in ssac/pkg/authz/interface.yaml.
func xqp30OwnerLookupQuery(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	have := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		have[q.Name] = true
	}
	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if d, ok := checkOwnershipMapping(p.File, om, have, seen); ok {
				diags = append(diags, d)
			}
		}
	}
	return diags
}
