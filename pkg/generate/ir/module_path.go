//ff:func feature=gen-ir type=util control=sequence
//ff:what modulePath -- manifest.backend.module 추출

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// modulePath extracts the Go module path from manifest.backend.module.
func modulePath(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil {
		return ""
	}
	return fs.Manifest.Backend.Module
}
