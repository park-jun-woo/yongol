//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-71 — @call cache.* requires manifest.cache.backend

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa71CacheBackendRequired validates XSA-71: if any SSaC service func
// calls cache.* built-ins, the manifest must declare cache.backend.
// See XSA-70 for the same rationale applied to sessions.
func xsa71CacheBackendRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if !usesCache(fs) {
		return nil
	}
	if fs.Manifest != nil && fs.Manifest.Cache != nil && fs.Manifest.Cache.Backend != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XSA-71] SSaC uses cache.* but manifest.cache.backend is not declared",
		Advice:  "Declare manifest.cache.backend (memory | redis)",
	}}
}
