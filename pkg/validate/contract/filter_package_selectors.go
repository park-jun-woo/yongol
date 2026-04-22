//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what filterPackageSelectors — recv 가 import 된 package 이름이면 제거 (DDL 필드 후보만 남김)

package contract

import "strings"

// filterPackageSelectors drops selectors whose receiver identifier
// matches a package imported by the same file. Example: with
// `import "database/sql"` the selector `sql.ErrNoRows` is kept out of
// DDLFields because `sql` is a package name, not a DDL model receiver.
// Pass a non-nil `pkgs` (empty is fine) — nil preserves every entry.
func filterPackageSelectors(selectors []string, pkgs map[string]bool) []string {
	if len(pkgs) == 0 {
		return selectors
	}
	out := make([]string, 0, len(selectors))
	for _, s := range selectors {
		idx := strings.Index(s, ".")
		if idx <= 0 {
			out = append(out, s)
			continue
		}
		if pkgs[s[:idx]] {
			continue
		}
		out = append(out, s)
	}
	return out
}
