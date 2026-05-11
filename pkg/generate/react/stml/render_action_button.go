//ff:func feature=stml-gen type=generator control=sequence
//ff:what Fields 없는 ActionBlock을 버튼 onClick JSX로 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderActionButton(a stmlparser.ActionBlock, indent int) string {
	ind := indentStr(indent)
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	tag := orDefault(a.Tag, "button")
	cls := clsAttr(a.ClassName)
	text := orDefault(a.SubmitText, a.OperationID)
	if tag == "button" {
		return fmt.Sprintf(`%s<button onClick={() => %s.mutate({})}%s>%s</button>`, ind, mutName, cls, text)
	}
	return fmt.Sprintf(`%s<%s%s><button onClick={() => %s.mutate({})}>%s</button></%s>`, ind, tag, cls, mutName, text, tag)
}
