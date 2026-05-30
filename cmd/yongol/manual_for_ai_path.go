//ff:func feature=cli type=util control=iteration dimension=1
//ff:what manualForAIPath — cwd 상위에서 manual-for-ai.md 를 탐지, 없으면 github URL 반환
package main

import (
	"os"
	"path/filepath"
)

// manualForAIPath walks up from the current working directory looking for
// manual-for-ai.md (present at the yongol module root). Returns its path when
// found, otherwise the canonical GitHub URL so the pointer is always resolvable.
func manualForAIPath() string {
	const githubURL = "https://github.com/park-jun-woo/yongol/blob/main/manual-for-ai.md"
	dir, err := os.Getwd()
	if err != nil {
		return githubURL
	}
	for {
		p := filepath.Join(dir, "manual-for-ai.md")
		if _, statErr := os.Stat(p); statErr == nil {
			return p
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return githubURL
		}
		dir = parent
	}
}
