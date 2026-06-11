//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what splitCaptureBinds — 캡처 바인딩을 token 필드 / refresh 필드 / claims 목록으로 분류

package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// splitCaptureBinds classifies the action's capture bindings for the
// commit renderer: the auth.token respField, the auth.refresh respField
// and the auth.claims.* bindings (plans/stml/sitemap Phase005). Sinks
// outside the whitelist contribute nothing (ParseCapture never yields
// them; the defensive result is no store write).
func splitCaptureBinds(captures []stmlparser.CaptureBind) (tokenField, refreshField string, claims []stmlparser.CaptureBind) {
	for _, c := range captures {
		switch c.Sink {
		case "auth.token":
			tokenField = c.RespField
		case "auth.refresh":
			refreshField = c.RespField
		default:
			claims = appendClaimBind(claims, c)
		}
	}
	return tokenField, refreshField, claims
}
