//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what frontendEnabled — manifest frontend 블록이 ON 상태인지 판정 (enabled!=false & 내용 있음)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// frontendEnabled reports whether the frontend is considered ON for this
// project. ProjectConfig.Frontend is a value type (not a pointer), so an
// omitted block still exists as a zero-value Frontend{}. ON is therefore
// decided by content signals rather than nil: the manifest must be present,
// frontend.enabled must not be explicitly false (nil == unspecified == ON),
// and the block must be meaningfully filled (Lang or Framework non-empty).
func frontendEnabled(fs *yongol.Fullstack) bool {
	if fs.Manifest == nil {
		return false
	}
	fe := fs.Manifest.Frontend
	if fe.Enabled != nil && !*fe.Enabled {
		return false
	}
	return fe.Lang != "" || fe.Framework != ""
}
