//ff:func feature=stml-gen type=generator control=sequence
//ff:what data-on-error 요소를 에러 상태 조건부 렌더 JSX로 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOnErrorElement renders the action's data-on-error slot: the element
// is shown only while the error state is set and its text binds the error
// message (e.g. {loginError && <p>{loginError}</p>}).
func renderOnErrorElement(se stmlparser.StaticElement, errVar string, indent int) string {
	ind := indentStr(indent)
	cls := clsAttr(se.ClassName)
	return fmt.Sprintf("%s{%s && <%s%s>{%s}</%s>}", ind, errVar, se.Tag, cls, errVar, se.Tag)
}
