//ff:func feature=stml-gen type=generator control=sequence
//ff:what data-on-error 미선언 액션의 기본 에러 표시 슬롯(role="alert")을 렌더링한다
package stml

import "fmt"

// renderDefaultOnErrorElement renders the default error slot for an action
// without a data-on-error element (page-flow Phase004 — silent-failure
// guard): a conditional <p role="alert"> bound to the action's error state,
// placed right after the submit button. The design token mirrors the field
// error styling (components_ui palette). Declaring data-on-error replaces
// this slot with the declared element at its declared position.
func renderDefaultOnErrorElement(errVar string, indent int) string {
	ind := indentStr(indent)
	return fmt.Sprintf(`%s{%s && <p role="alert" className="text-sm text-destructive">{%s}</p>}`, ind, errVar, errVar)
}
