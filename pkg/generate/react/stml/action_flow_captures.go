//ff:func feature=stml-gen type=util control=sequence
//ff:what bearer 모드에서만 ActionBlock의 data-capture 바인딩을 유효 캡처로 반환한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// actionFlowCaptures returns the action's effective data-capture bindings.
// Captures commit to the bearer session store, so they are emitted only in
// bearer mode; in cookie mode (httpOnly cookies — nothing to capture) the
// declarations are ignored here and TM-24 diagnoses them at validate time.
func actionFlowCaptures(a stmlparser.ActionBlock, bearerAuth bool) []stmlparser.CaptureBind {
	if !bearerAuth {
		return nil
	}
	return a.Captures
}
