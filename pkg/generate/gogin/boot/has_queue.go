//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasQueue — manifest.queue.backend 존재 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasQueue returns true when the manifest declares queue.backend.
func hasQueue(fs *yongol.Fullstack) bool {
	return fs.Manifest != nil && fs.Manifest.Queue != nil && fs.Manifest.Queue.Backend != ""
}
