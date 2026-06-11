//ff:func feature=stml-gen type=generator control=sequence
//ff:what data-bind 필드의 JSX를 스키마 타입에 따라 생성한다 (img는 src 바인딩, 그 외 타입별 children)
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderBindJSX generates JSX for a data-bind field. An <img> tag binds the
// value to the src attribute (the tag declares media intent — SSOT decision),
// emitting a self-closing element with no children. Every other tag wraps the
// type-aware value expression (bindValueExpr) as children. With a zero bindCtx
// the value falls back to {path}, reproducing the pre-Phase037 output exactly.
func renderBindJSX(b stmlparser.FieldBind, dataVar string, indent int, ctx bindCtx) string {
	ind := indentStr(indent)
	cls := clsAttr(b.ClassName)
	path := dataVar + "." + optionalChainPath(b.Name)

	if b.Tag == "img" {
		return fmt.Sprintf("%s<img src={%s} alt=%q%s />", ind, path, toLabel(b.Name), cls)
	}

	tag := orDefault(b.Tag, "span")
	return fmt.Sprintf("%s<%s%s>%s</%s>", ind, tag, cls, bindValueExpr(path, ctx.field(b.Name)), tag)
}
