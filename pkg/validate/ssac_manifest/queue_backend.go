//ff:func feature=validate type=util control=sequence topic=config-check
//ff:what manifest queue.backend 값 반환 (unset 이면 "")

package ssac_manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// queueBackend returns the manifest queue.backend value ("postgres"/"memory"/...)
// or "" if unset. Ground's Config map only flags key presence, so we read the
// manifest struct directly.
func queueBackend(fs *yongol.Fullstack) string {
	if fs.Manifest == nil || fs.Manifest.Queue == nil {
		return ""
	}
	return fs.Manifest.Queue.Backend
}
