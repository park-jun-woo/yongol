//ff:func feature=orchestrator type=util control=sequence
//ff:what resolveDesignPath — manifest 또는 convention 기반으로 DESIGN.md 경로 결정

package yongol

import "path/filepath"

// resolveDesignPath determines the DESIGN.md path. Priority:
// 1. manifest.frontend.design (explicit path relative to specs root)
// 2. convention: frontend/DESIGN.md
func resolveDesignPath(fs *Fullstack, root string) string {
	if fs.Manifest != nil && fs.Manifest.Frontend.Design != "" {
		return filepath.Join(root, fs.Manifest.Frontend.Design)
	}
	// Convention fallback.
	return filepath.Join(root, "frontend", "DESIGN.md")
}
