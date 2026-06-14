//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what actionHasRefreshCapture — 한 액션의 data-capture 중 auth.refresh sink가 있는지 판정

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// actionHasRefreshCapture reports whether an action declares an auth.refresh
// data-capture. The sink string mirrors the parser whitelist
// (splitCaptureBinds / type_capture_bind.go). Extracted from
// hasRefreshCaptures to keep that page-walk within the nesting limit.
func actionHasRefreshCapture(a stml.ActionBlock) bool {
	for _, c := range a.Captures {
		if c.Sink == "auth.refresh" {
			return true
		}
	}
	return false
}
