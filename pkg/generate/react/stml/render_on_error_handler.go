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
// object, so `err` is usually not an Error instance. `message` is read
// defensively: XOE-01 only checks error/code required (WARNING) and never
// `message`, so the schema gives no type/presence guarantee — extract a
// non-empty string `message` first, fall back to String(err).
func renderOnErrorHandler(a stmlparser.ActionBlock) string {
	errVar := errorStateVar(a)
	return fmt.Sprintf(`    onError: (err) => {
      const msg = (err as any)?.message
      set%s(typeof msg === 'string' && msg !== '' ? msg : String(err))
    },
`, toUpperFirst(errVar))
}
