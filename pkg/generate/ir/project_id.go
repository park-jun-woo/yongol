//ff:func feature=gen-ir type=util control=sequence
//ff:what projectID -- manifest.metadata.name 추출

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// projectID extracts the project name from manifest.metadata.name.
func projectID(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil {
		return ""
	}
	return fs.Manifest.Metadata.Name
}
