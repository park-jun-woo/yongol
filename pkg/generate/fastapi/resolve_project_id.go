//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveProjectID — manifest.metadata.name → 프로젝트 ID 해석

package fastapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveProjectID extracts the project name from the manifest, falling
// back to "app" when not available.
func resolveProjectID(fs *yongol.Fullstack) string {
	if fs.Manifest != nil && fs.Manifest.Metadata.Name != "" {
		return fs.Manifest.Metadata.Name
	}
	return "app"
}
