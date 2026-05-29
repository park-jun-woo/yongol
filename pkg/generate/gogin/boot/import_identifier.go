//ff:func feature=gen-gogin type=util control=sequence
//ff:what importIdentifier — import 라인에서 패키지 식별자 (마지막 경로 세그먼트) 추출

package boot

import "strings"

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
	i := strings.Index(s, "\"")
	if i < 0 {
		return ""
	}
	s = s[i:]
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
