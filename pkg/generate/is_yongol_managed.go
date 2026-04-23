//ff:func feature=generate type=util control=selection
//ff:what isYongolManaged — specs/frontend 상대경로가 생성기 관할(api.ts, types/, lib/, components/ui/) 인지 판별
package generate

import (
	"path/filepath"
	"strings"
)

// isYongolManaged is true when the relative path (inside specs/frontend/)
// names a subtree that the generator owns and should never be overwritten
// from specs. Users are free to author there but any conflict is resolved
// in favor of the generator.
func isYongolManaged(rel string) bool {
	rel = filepath.ToSlash(rel)
	switch {
	case rel == "src/api.ts",
		strings.HasPrefix(rel, "src/types/"),
		strings.HasPrefix(rel, "src/lib/"),
		strings.HasPrefix(rel, "src/components/ui/"):
		return true
	}
	return false
}
