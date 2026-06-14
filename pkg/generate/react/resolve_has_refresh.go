//ff:func feature=gen-react type=util control=sequence
//ff:what resolveHasRefresh — bearer store가 refresh 토큰 필드를 방출해야 하는지 판정 (manifest refresh_field ∨ STML auth.refresh 캡처)

package react

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveHasRefresh reports whether the bearer session store must carry the
// refresh token — the `refresh` field, setAuth's second argument and clear's
// reset of it (BUG-135). It is the single signal that keeps three emitters in
// type agreement so the store never drops a field another emitter still
// writes or reads:
//   - frontend.auth.refresh_field declared → resolveRefreshPlan's
//     precondition, the api.ts 401 refresh flow reads store.refresh; and
//   - an STML auth.refresh data-capture exists → renderCaptureCommit writes
//     it via setAuth(token, refresh).
//
// When neither holds (the common bearer-without-refresh project) the refresh
// surface is dead code and is dropped.
func resolveHasRefresh(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	if fs.Manifest != nil && fs.Manifest.Frontend.Auth != nil && fs.Manifest.Frontend.Auth.RefreshField != "" {
		return true
	}
	return hasRefreshCaptures(fs.STMLPages)
}
