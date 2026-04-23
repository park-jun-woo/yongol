//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-75 — manifest.cache.backend declared but SSaC never calls cache.*

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa75CacheBackendUnused validates XSA-75: the manifest declares a cache
// backend but no SSaC service func calls cache.*. See XSA-74 for the same
// rationale applied to sessions.
func xsa75CacheBackendUnused(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if fs.Manifest == nil || fs.Manifest.Cache == nil || fs.Manifest.Cache.Backend == "" {
		return nil
	}
	if usesCache(fs) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[XSA-75] manifest.cache.backend is declared but no SSaC function uses cache.*",
		Advice:  "Remove manifest.cache.backend or add an @call cache.* sequence",
	}}
}
