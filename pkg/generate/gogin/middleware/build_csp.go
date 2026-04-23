//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=security-headers
//ff:what buildCSPValue — directives map → "key1 src1 src2; key2 src3" 형식의 CSP 문자열 조합

package middleware

import (
	"sort"
	"strings"
)

// buildCSPValue mirrors the runtime BuildCSPValue function emitted into the
// generated security_headers.go file. Kept as a separate Go symbol so unit
// tests in this package can exercise the algorithm directly — the runtime
// helper lives inside a const string template and cannot be invoked from
// yongol's own tests. Any change here MUST be matched inside
// securityHeadersSource to keep generator-side and runtime-side behavior
// aligned.
//
// Contract:
//   - Empty / nil map → empty string (caller suppresses header).
//   - Keys emitted in sorted order for reproducible output.
//   - A directive with no sources emits just the directive name (legal CSP).
func buildCSPValue(directives map[string][]string) string {
	if len(directives) == 0 {
		return ""
	}
	keys := make([]string, 0, len(directives))
	for k := range directives {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		sources := directives[k]
		if len(sources) == 0 {
			parts = append(parts, k)
			continue
		}
		parts = append(parts, k+" "+strings.Join(sources, " "))
	}
	return strings.Join(parts, "; ")
}
