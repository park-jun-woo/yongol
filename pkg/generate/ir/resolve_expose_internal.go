//ff:func feature=gen-ir type=util control=sequence
//ff:what resolveExposeInternal -- manifest.backend.error.expose_internal_error 기본값 해석

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveExposeInternal reads manifest.backend.error.expose_internal_error.
func resolveExposeInternal(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	e := fs.Manifest.Backend.Error
	if e == nil || e.ExposeInternalError == nil {
		return false
	}
	return *e.ExposeInternalError
}
