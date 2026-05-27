//ff:func feature=gen-ir type=util control=sequence
//ff:what buildCORSConfig -- manifest CORS 설정 → CORSBootConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildCORSConfig extracts CORS configuration from the manifest and returns
// a CORSBootConfig. Returns nil when CORS is not configured.
func buildCORSConfig(fs *yongol.Fullstack) *CORSBootConfig {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.CORS == nil {
		return nil
	}
	c := fs.Manifest.Backend.CORS
	return &CORSBootConfig{
		AllowOrigins:     c.AllowOrigins,
		AllowCredentials: c.AllowCredentials,
	}
}
