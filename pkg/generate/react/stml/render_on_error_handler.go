//ff:func feature=stml-gen type=generator control=sequence
//ff:what 모든 액션의 mutation onError 핸들러(에러 메시지 상태 갱신)를 렌더링한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOnErrorHandler renders the mutation onError handler that feeds the
// action's error state. Always emitted since page-flow Phase004 — without a
// data-on-error declaration the state feeds the default error slot instead,
// so a rejected mutation never fails silently (BUG-113 (2)).
//
// Since BUG-113 the api wrapper throws the server ErrorResponse as a plain
// object, so `err` is usually not an Error instance. displayField is the
// schema-derived display field (ExtractErrorDisplayField: "error" for the
// current ErrorResponse contract) and is read first; `message` is read second
// to keep covering real Error instances from network-level rejects
// (BUG-125). String(err) is the final fallback only when both are absent or
// non-string. An empty displayField is normalized to "error" so partially
// constructed GenerateOptions (no ErrorDisplayField) never emit a broken
// `e?. ?? e?.message` expression.
func renderOnErrorHandler(a stmlparser.ActionBlock, displayField string) string {
	if displayField == "" {
		displayField = "error"
	}
	errVar := errorStateVar(a)
	return fmt.Sprintf(`    onError: (err) => {
      const e = err as any
      const msg = e?.%s ?? e?.message
      set%s(typeof msg === 'string' && msg !== '' ? msg : String(err))
    },
`, displayField, toUpperFirst(errVar))
}
