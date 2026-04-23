//ff:func feature=gen-gogin type=util control=sequence topic=error-envelope
//ff:what resolveExposeInternalError — expose_internal_error 컴파일타임 기본값 (기본 false)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveExposeInternalError returns the compile-time default for the
// ExposeInternalError flag. Defaults to false (production-safe). Reads
// manifest.backend.error.expose_internal_error when present.
func resolveExposeInternalError(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	e := fs.Manifest.Backend.Error
	if e == nil || e.ExposeInternalError == nil {
		return false
	}
	return *e.ExposeInternalError
}
