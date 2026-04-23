//ff:func feature=generate type=util control=sequence
//ff:what manifestDeclaresFile — manifest.file.backend 선언 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// manifestDeclaresFile returns true when the manifest explicitly
// declares file.backend. Nil-safe on every hop.
func manifestDeclaresFile(fs *yongol.Fullstack) bool {
	return fs != nil && fs.Manifest != nil &&
		fs.Manifest.File != nil && fs.Manifest.File.Backend != ""
}
