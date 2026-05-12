//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what defaultLayoutFromManifest — manifest 에서 defaultLayout 값 추출

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// defaultLayoutFromManifest extracts the defaultLayout value from the manifest,
// returning empty string if the manifest is nil.
func defaultLayoutFromManifest(fs *yongol.Fullstack) string {
	if fs.Manifest == nil {
		return ""
	}
	return fs.Manifest.Frontend.DefaultLayout
}
