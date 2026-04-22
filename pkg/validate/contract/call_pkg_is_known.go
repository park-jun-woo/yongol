//ff:func feature=validate-contract type=util control=sequence
//ff:what callPkgIsKnown — "pkg.Func" 의 pkg 접두사가 SSOT 알려진 패키지 집합에 속하는지 검사

package contract

import "strings"

// callPkgIsKnown returns true when the package prefix of `pkg.Func` is
// tracked by the SSOT. Calls whose receiver is unknown (std library,
// out-of-band helper) are silently ignored by PRV-02 — flagging them
// would produce noise with no actionable fix.
func callPkgIsKnown(call string, knownPkgs map[string]bool) bool {
	idx := strings.Index(call, ".")
	if idx <= 0 {
		return false
	}
	return knownPkgs[call[:idx]]
}
