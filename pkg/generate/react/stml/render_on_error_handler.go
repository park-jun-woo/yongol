//ff:func feature=stml-gen type=generator control=sequence
//ff:what data-on-error 선언 시 mutation onError 핸들러(에러 메시지 상태 갱신)를 렌더링한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOnErrorHandler renders the mutation onError handler that feeds the
// action's data-on-error slot. Empty when the action declares no
// data-on-error element (no handler is emitted — current behavior kept).
func renderOnErrorHandler(a stmlparser.ActionBlock) string {
	errVar := errorStateVar(a)
	if errVar == "" {
		return ""
	}
	return fmt.Sprintf(`    onError: (err) => {
      set%s(err instanceof Error ? err.message : String(err))
    },
`, toUpperFirst(errVar))
}
