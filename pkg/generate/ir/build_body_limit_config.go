//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildBodyLimitConfig -- manifest.backend.http → BodyLimitConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildBodyLimitConfig always returns a config (body-limit is always active).
func buildBodyLimitConfig(fs *yongol.Fullstack) *BodyLimitConfig {
	cfg := &BodyLimitConfig{
		BodyLimit:          1048576,  // 1 MiB default
		MultipartLimit:     33554432, // 32 MiB default
		BodyOverrides:      map[string]int64{},
		MultipartOverrides: map[string]int64{},
	}
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return cfg
	}
	h := fs.Manifest.Backend.HTTP
	if h.BodyLimit != "" {
		cfg.BodyLimit = parseBodyLimitOrDefault(h.BodyLimit, cfg.BodyLimit)
	}
	if h.MultipartLimit != "" {
		cfg.MultipartLimit = parseBodyLimitOrDefault(h.MultipartLimit, cfg.MultipartLimit)
	}
	return cfg
}
