//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolveHeaderLimit — manifest.backend.http.header_limit 값 결정 (파싱 실패 시 1 MiB)

package boot

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/middleware"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveHeaderLimit returns the manifest-derived MaxHeaderBytes default.
// Nil manifest / empty value / parse failure → 1 MiB (Go stdlib default).
func resolveHeaderLimit(fs *yongol.Fullstack) int64 {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return defaultHeaderLimit
	}
	raw := fs.Manifest.Backend.HTTP.HeaderLimit
	if raw == "" {
		return defaultHeaderLimit
	}
	n, err := middleware.ParseSize(raw)
	if err != nil {
		return defaultHeaderLimit
	}
	return n
}
