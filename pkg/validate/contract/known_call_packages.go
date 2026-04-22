//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what knownCallPackages — expCalls 키의 "pkg" 접두사 집합 (SSOT 가 추적하는 패키지)

package contract

import "strings"

// knownCallPackages extracts the set of package identifiers that the
// SSOT's Func / SSaC surface declares. PRV-02 only flags missing
// `pkg.Func` calls when pkg belongs to this set — other callers
// (std library, internal helpers unrelated to SSOT) are out of
// scope and would produce noise rather than actionable drift.
func knownCallPackages(expCalls map[string]bool) map[string]bool {
	out := map[string]bool{}
	for k := range expCalls {
		idx := strings.Index(k, ".")
		if idx <= 0 {
			continue
		}
		out[k[:idx]] = true
	}
	return out
}
