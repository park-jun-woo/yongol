//ff:func feature=gen-react type=util control=sequence
//ff:what resolveTheme — Fullstack.Manifest.Frontend.Theme 조회 (없으면 nil)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveTheme extracts the Frontend.Theme pointer (may be nil).
func resolveTheme(fs *yongol.Fullstack) *manifestTheme {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	return fs.Manifest.Frontend.Theme
}
