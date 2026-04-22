//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasFile — manifest.file.backend 존재 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasFile returns true when the manifest declares file.backend.
func hasFile(fs *yongol.Fullstack) bool {
	return fs.Manifest != nil && fs.Manifest.File != nil && fs.Manifest.File.Backend != ""
}
