//ff:func feature=gen-ir type=util control=sequence
//ff:what corsIsEnabled -- manifest.backend.cors.enabled 여부 판정

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// corsIsEnabled returns true when manifest.backend.cors.enabled is set.
func corsIsEnabled(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	c := fs.Manifest.Backend.CORS
	return c != nil && c.Enabled
}
