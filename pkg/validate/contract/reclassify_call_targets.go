//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what reclassifyCallTargets — "recv.Func" 중 recv가 import 패키지 아니면 query 로 재분류

package contract

import "strings"

// reclassifyCallTargets splits the raw CallTargets list into two
// buckets depending on whether the receiver identifier is a known
// imported package:
//
//   - External pkg.Func calls stay in pkgCalls and are compared
//     against the Func SSOT / SSaC @call expected set.
//   - local.Method calls (e.g. `qtx.UserFindByID`) move to localMethods
//     and are compared against the sqlc query expected set — these are
//     the common pattern for `server.Queries.WithTx(tx)` followed by
//     methods on the resulting `qtx` local.
//
// Calls without a dot and calls on the denylist are dropped
// (defensive — the contract extractor already filters those).
func reclassifyCallTargets(calls []string, pkgs map[string]bool) (pkgCalls, localMethods []string) {
	for _, c := range calls {
		idx := strings.Index(c, ".")
		if idx <= 0 {
			continue
		}
		recv := c[:idx]
		method := c[idx+1:]
		if pkgs[recv] {
			pkgCalls = append(pkgCalls, c)
			continue
		}
		localMethods = append(localMethods, method)
	}
	return pkgCalls, localMethods
}
