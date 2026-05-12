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
	if tag == "button" {
		return fmt.Sprintf(`%s<Button%s onClick={() => %s.mutate(%s)} disabled={%s.isPending}%s>%s</Button>`, ind, variant, mutName, mutateArg, mutName, cls, pendingExpr)
	}
	return fmt.Sprintf(`%s<%s%s><Button%s onClick={() => %s.mutate(%s)} disabled={%s.isPending}>%s</Button></%s>`, ind, tag, cls, variant, mutName, mutateArg, mutName, pendingExpr, tag)
}
