//ff:func feature=validate type=util control=sequence topic=tsx
//ff:what TSX @/ path alias 를 실제 디렉토리로 해석

package tsx

import (
	"os"
	"path/filepath"
)

// resolveAliasRoot maps the @/ path-alias to a concrete directory.
// Convention:
//   - <frontend>/src exists  → @/ = <frontend>/src
//   - otherwise               → @/ = <frontend>
//
// This matches the Vite/tsconfig default where pages live under src/ but
// keeps legacy layouts (components directly under frontend/) working.
func resolveAliasRoot(frontendRoot string) string {
	if st, err := os.Stat(filepath.Join(frontendRoot, "src")); err == nil && st.IsDir() {
		return filepath.Join(frontendRoot, "src")
	}
	return frontendRoot
}
