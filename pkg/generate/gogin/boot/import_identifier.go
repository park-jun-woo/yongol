//ff:func feature=gen-gogin type=util control=sequence
//ff:what importIdentifier — import 라인에서 패키지 식별자 (마지막 경로 세그먼트) 추출

package boot

import (
	"strings"
	"unicode"
)

// importIdentifier extracts the package identifier from an import line
// such as `"strconv"` → "strconv" or `_ "github.com/lib/pq"` → "pq".
//
// Explicit aliases win when present — `limiter "github.com/ulule/limiter/v3"`
// returns "limiter". Bare paths whose last segment is a SemVer-major
// marker (`/v2`, `/v3`, …) fall through to the prior segment since Go
// treats the path as the package's parent directory name.
func importIdentifier(imp string) string {
	s := strings.TrimSpace(imp)
	// Detect explicit alias: `alias "path"` or `_ "path"`.
	if q := strings.Index(s, "\""); q > 0 {
		alias := strings.TrimSpace(s[:q])
		if alias != "" && alias != "_" {
			return alias
		}
	}
	if i := strings.Index(s, "\""); i >= 0 {
		s = s[i:]
	}
	s = strings.Trim(s, "\"")
	if s == "" {
		return ""
	}
	parts := strings.Split(s, "/")
	last := parts[len(parts)-1]
	if isSemVerMajor(last) && len(parts) >= 2 {
		return parts[len(parts)-2]
	}
	return last
}

// isSemVerMajor reports whether s looks like "v2", "v3", … as used by Go
// module versioned path suffixes. Pure numeric suffix matters — "v3alpha"
// is a real package name.
func isSemVerMajor(s string) bool {
	if len(s) < 2 || (s[0] != 'v' && s[0] != 'V') {
		return false
	}
	for _, r := range s[1:] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
