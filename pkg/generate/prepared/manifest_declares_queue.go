//ff:func feature=generate type=util control=sequence
//ff:what manifestDeclaresQueue — manifest.queue.backend 선언 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// manifestDeclaresQueue returns true when the manifest explicitly
// declares queue.backend. Nil-safe on every hop.
func manifestDeclaresQueue(fs *yongol.Fullstack) bool {
	return fs != nil && fs.Manifest != nil &&
		fs.Manifest.Queue != nil && fs.Manifest.Queue.Backend != ""
}
