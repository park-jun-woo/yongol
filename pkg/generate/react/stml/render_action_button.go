//ff:func feature=stml-gen type=generator control=sequence
//ff:what Fields 없는 ActionBlock을 Button onClick JSX로 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderActionButton(a stmlparser.ActionBlock, indent int, noBodyOps map[string]bool) string {
	ind := indentStr(indent)
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	tag := orDefault(a.Tag, "button")
	cls := clsAttr(a.ClassName)
	text := orDefault(a.SubmitText, a.OperationID)
	pendingExpr := fmt.Sprintf("{%s.isPending ? '처리 중...' : '%s'}", mutName, text)
	variant := ""
	if isDeleteOperation(a.OperationID) {
		variant = ` variant="destructive"`
	}
	mutateArg := "{}"
	if noBodyOps[a.OperationID] {
		mutateArg = ""
	}
	disabledExpr := mergeDisabledExpr(mutName+".isPending", a.EnabledWhen, "data")
	// data-on-error: the button path drops static children, so the error
	// slot is rendered right after the button (conditional on the state).
	errJSX := ""
	if errVar := errorStateVar(a); errVar != "" {
		se := findOnErrorStatic(a.Children)
		if se == nil {
			se = &stmlparser.StaticElement{Tag: "p"}
		}
		errJSX = "\n" + renderOnErrorElement(*se, errVar, indent)
	}
	if tag == "button" {
		return fmt.Sprintf(`%s<Button%s onClick={() => %s.mutate(%s)} disabled={%s}%s>%s</Button>%s`, ind, variant, mutName, mutateArg, disabledExpr, cls, pendingExpr, errJSX)
	}
	return fmt.Sprintf(`%s<%s%s><Button%s onClick={() => %s.mutate(%s)} disabled={%s}>%s</Button></%s>%s`, ind, tag, cls, variant, mutName, mutateArg, disabledExpr, pendingExpr, tag, errJSX)
}
