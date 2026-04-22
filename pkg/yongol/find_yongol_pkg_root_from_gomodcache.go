//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what findYongolPkgRootFromGoModCache — $GOMODCACHE 의 최신 ssac@<ver>/pkg 경로 반환

package yongol

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// findYongolPkgRootFromGoModCache searches
// `$GOMODCACHE/github.com/park-jun-woo/ssac@<version>/pkg` and returns the
// latest semver-sorted directory that exists. Used when yongol is invoked
// from a project that depends on ssac via Go modules rather than as a sibling
// clone. Returns "" when no candidate exists.
func findYongolPkgRootFromGoModCache() string {
	cache := goModCache()
	if cache == "" {
		return ""
	}
	root := filepath.Join(cache, "github.com", "park-jun-woo")
	entries, err := os.ReadDir(root)
	if err != nil {
		return ""
	}
	var candidates []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() && strings.HasPrefix(name, "ssac@") {
			candidates = append(candidates, name)
		}
	}
	if len(candidates) == 0 {
		return ""
	}
	sort.Sort(sort.Reverse(sort.StringSlice(candidates)))
	for _, c := range candidates {
		candidate := filepath.Join(root, c, "pkg")
		if isDir(candidate) {
			return candidate
		}
	}
	return ""
}
