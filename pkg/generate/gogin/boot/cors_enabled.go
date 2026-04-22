//ff:func feature=gen-gogin type=util control=sequence
//ff:what corsEnabled — manifest.backend.cors.enabled 여부 판정

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// corsEnabled returns true when manifest declares backend.cors.enabled.
func corsEnabled(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	c := fs.Manifest.Backend.CORS
	return c != nil && c.Enabled
}
