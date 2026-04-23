//ff:func feature=tsx-parser type=util control=sequence
//ff:what isLocalComponentImport — `@/components/` / `./components/` / sibling 경로 감지

package tsx

import "strings"

// isLocalComponentImport matches the two conventions emitted by yongol's
// React scaffold: path-alias `@/components/...` and relative `./components/...`
// (or `../components/...`). Anything else — npm packages, absolute HTTP URLs,
// deep path-aliases like `@/lib/...` — is skipped.
func isLocalComponentImport(src string) bool {
	if src == "" {
		return false
	}
	if strings.HasPrefix(src, "@/components/") {
		return true
	}
	if strings.HasPrefix(src, "./components/") || strings.HasPrefix(src, "../components/") {
		return true
	}
	// Relative sibling paths inside pages/ that still live under components/.
	if strings.HasPrefix(src, "./") || strings.HasPrefix(src, "../") {
		if strings.Contains(src, "/components/") {
			return true
		}
	}
	return false
}
