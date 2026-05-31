//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what lookupOwnershipIR -- Rego @ownership 어노테이션에서 resource 매칭 OwnershipInfo 조회

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// lookupOwnershipIR returns ownership lookup metadata for the given resource
// from the Fullstack's parsed Rego @ownership annotations. Returns nil when no
// matching ownership annotation is found.
func lookupOwnershipIR(fs *yongol.Fullstack, resource string) *OwnershipInfo {
	for _, p := range fs.ParsedPolicies {
		if info := matchOwnershipIR(fs, p.Ownerships, resource); info != nil {
			return info
		}
	}
	return nil
}
