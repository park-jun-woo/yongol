//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what backendAuthMode — backend.auth가 선언된 경우 ResolvedMode, 없으면 빈 문자열

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// backendAuthMode returns the effective backend.auth mode ("cookie",
// "bearer", or "hybrid") or "" when no backend.auth block is declared at
// all. Rules gate on the returned value so a project without auth stays
// out of the TM-21/22/24 flow rules entirely.
func backendAuthMode(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return ""
	}
	return fs.Manifest.Backend.Auth.ResolvedMode()
}
