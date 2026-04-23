//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx
//ff:what TSX import 대상이 .tsx/.ts/.jsx/.js/index 로 해석되는지 검사

package tsx

import "os"

// fileExistsWithTsxVariants returns true if any common resolution of `base`
// exists: base.tsx, base.ts, base.jsx, base.js, base/index.tsx, base/index.ts.
// The order mirrors Node's moduleResolution with TypeScript+JSX.
func fileExistsWithTsxVariants(base string) bool {
	// Direct file match (with extension already present).
	if _, err := os.Stat(base); err == nil {
		return true
	}
	candidates := tsxResolveCandidates(base)
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return true
		}
	}
	return false
}
