//ff:func feature=stml-gen type=generator control=sequence
//ff:what data-bind 필드의 JSX를 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderBindJSX generates JSX for a data-bind field.
func renderBindJSX(b stmlparser.FieldBind, dataVar string, indent int) string {
	ind := indentStr(indent)
	tag := orDefault(b.Tag, "span")
	cls := clsAttr(b.ClassName)
	return fmt.Sprintf("%s<%s%s>{%s.%s}</%s>", ind, tag, cls, dataVar, b.Name, tag)
}
