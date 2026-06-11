//ff:func feature=stml-gen type=generator control=sequence
//ff:what FetchBlock의 로딩/에러/데이터 조건부 JSX를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFetchJSX generates JSX for a FetchBlock using ChildNode tree. The
// bindCtx scope is re-set to this fetch's operationId so descendant binds —
// including nested data-fetch blocks dispatched via renderChildNodes — resolve
// their field types against the correct op (plans/gen/frontend Phase037).
func renderFetchJSX(f stmlparser.FetchBlock, indent int, noBodyOps map[string]bool, ctx bindCtx) string {
	ctx.opID = f.OperationID
	alias := toLowerFirst(f.OperationID) + "Data"
	ind := indentStr(indent)
	tag := orDefault(f.Tag, "div")
	cls := clsAttr(f.ClassName)

	var lines []string
	lines = append(lines, fmt.Sprintf("%s{%sLoading && <div>로딩 중...</div>}", ind, alias))
	lines = append(lines, fmt.Sprintf("%s{%sError && <div>오류가 발생했습니다</div>}", ind, alias))
	lines = append(lines, fmt.Sprintf("%s{%s && (", ind, alias))
	lines = append(lines, fmt.Sprintf("%s  <%s%s>", ind, tag, cls))

	lines = append(lines, renderFetchJSXBody(f, alias, indent+4, noBodyOps, ctx)...)

	lines = append(lines, fmt.Sprintf("%s  </%s>", ind, tag))
	lines = append(lines, fmt.Sprintf("%s)}", ind))

	return strings.Join(lines, "\n")
}
