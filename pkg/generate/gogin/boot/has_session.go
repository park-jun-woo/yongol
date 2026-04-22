//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasSession — manifest.session.backend 존재 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasSession returns true when the manifest declares session.backend.
func hasSession(fs *yongol.Fullstack) bool {
	return fs.Manifest != nil && fs.Manifest.Session != nil && fs.Manifest.Session.Backend != ""
}
