//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what findYongolPkgRootFromCWD — CWD 상위에서 yongol root 찾고 형제 ssac/pkg 반환

package yongol

import (
	"os"
	"path/filepath"
)

// findYongolPkgRootFromCWD walks up from the current working directory until
// it locates a yongol repo root, then returns the sibling `ssac/pkg`
// directory via trySSaCPkgPath. Used for development worktrees where yongol
// and ssac live as sibling clones.
func findYongolPkgRootFromCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		if isYongolRoot(dir) {
			return trySSaCPkgPath(dir)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
