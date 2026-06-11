//ff:func feature=stml-gen type=generator control=sequence
//ff:what Fields 있는 ActionBlock을 form JSX로 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderActionForm(a stmlparser.ActionBlock, indent int) string {
	ind := indentStr(indent)
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	formName := toLowerFirst(a.OperationID) + "Form"
	idPrefix := toLowerFirst(a.OperationID)
	cls := clsAttr(a.ClassName)
	submitText := orDefault(a.SubmitText, "제출")

	var lines []string
	lines = append(lines, fmt.Sprintf(`%s<form onSubmit={%s.handleSubmit((data) => %s.mutate(data))}%s>`, ind, formName, mutName, cls))

	if len(a.Children) > 0 {
		lines = append(lines, renderActionChildNodes(a.Children, formName, idPrefix, errorStateVar(a), indent+2)...)
	} else {
		for _, f := range a.Fields {
			lines = append(lines, renderFieldJSX(f, formName, idPrefix, indent+2))
		}
	}

	disabledExpr := mergeDisabledExpr(mutName+".isPending", a.EnabledWhen, "data")
	lines = append(lines, fmt.Sprintf(`%s  <Button type="submit" disabled={%s}>{%s.isPending ? '처리 중...' : '%s'}</Button>`, ind, disabledExpr, mutName, submitText))
	// page-flow Phase004: no data-on-error declared — emit the default error
	// slot right after the submit button so a rejected mutation stays visible.
	if !a.OnErrorNode {
		lines = append(lines, renderDefaultOnErrorElement(errorStateVar(a), indent+2))
	}
	lines = append(lines, fmt.Sprintf(`%s</form>`, ind))

	return strings.Join(lines, "\n")
}
