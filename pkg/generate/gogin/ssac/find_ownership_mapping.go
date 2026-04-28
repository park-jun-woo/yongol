//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what findOwnershipMapping — @ownership 매핑에서 seq.Resource 에 해당하는 첫 항목을 반환

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/rego"

// findOwnershipMapping returns the first parsed `@ownership` mapping whose
// Resource matches the @auth sequence's Resource. Nil when no mapping
// applies — the caller then omits the lookup.
func findOwnershipMapping(ownerships []rego.OwnershipMapping, resource string) *rego.OwnershipMapping {
	for i := range ownerships {
		if ownerships[i].Resource == resource {
			return &ownerships[i]
		}
	}
	return nil
}
