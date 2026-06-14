//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what hasRefreshCaptures — 어느 페이지든 auth.refresh data-capture가 선언되었는지 판정

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// hasRefreshCaptures reports whether any page action declares an auth.refresh
// data-capture. When one exists renderCaptureCommit emits a two-arg
// setAuth(token, refresh), so the session store must keep the refresh field
// to stay type-consistent with the capture commit (BUG-135 store/capture
// alignment). The sink string mirrors the parser whitelist (splitCaptureBinds
// / type_capture_bind.go).
func hasRefreshCaptures(pages []stml.PageSpec) bool {
	for _, p := range pages {
		for _, a := range p.Actions {
			if actionHasRefreshCapture(a) {
				return true
			}
		}
	}
	return false
}
