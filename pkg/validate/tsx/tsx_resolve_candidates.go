//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx
//ff:what TSX/JSX 확장자 + index 모듈 해석 후보 경로 목록 생성

package tsx

import "path/filepath"

// tsxResolveCandidates returns filesystem paths that represent all the
// common TypeScript/JSX module resolutions of `base` — both direct
// extension variants and directory-index variants.
func tsxResolveCandidates(base string) []string {
	suffixes := []string{".tsx", ".ts", ".jsx", ".js"}
	indexes := []string{"index.tsx", "index.ts", "index.jsx", "index.js"}
	out := make([]string, 0, len(suffixes)+len(indexes))
	for _, s := range suffixes {
		out = append(out, base+s)
	}
	for _, idx := range indexes {
		out = append(out, filepath.Join(base, idx))
	}
	return out
}
