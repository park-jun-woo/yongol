//ff:func feature=orchestrator type=util control=sequence
//ff:what goModCache — Go module cache 디렉토리 경로 해석

package yongol

import (
	"os"
	"path/filepath"
)

// goModCache returns the resolved Go module cache directory, honoring
// `GOMODCACHE`, falling back to `$GOPATH/pkg/mod`, then `$HOME/go/pkg/mod`.
// Returns "" when no candidate is resolvable.
func goModCache() string {
	if v := os.Getenv("GOMODCACHE"); v != "" {
		return v
	}
	if v := os.Getenv("GOPATH"); v != "" {
		return filepath.Join(v, "pkg", "mod")
	}
	if home, err := os.UserHomeDir(); err == nil {
		return filepath.Join(home, "go", "pkg", "mod")
	}
	return ""
}
