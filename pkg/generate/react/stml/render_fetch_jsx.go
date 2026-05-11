//ff:func feature=stml-gen type=generator control=sequence
//ff:what FetchBlock의 로딩/에러/데이터 조건부 JSX를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFetchJSX generates JSX for a FetchBlock using ChildNode tree.
func renderFetchJSX(f stmlparser.FetchBlock, indent int) string {
	alias := toLowerFirst(f.OperationID) + "Data"
	ind := indentStr(indent)
	tag := orDefault(f.Tag, "div")
	cls := clsAttr(f.ClassName)

	var lines []string
	lines = append(lines, fmt.Sprintf("%s{%sLoading && <div>로딩 중...</div>}", ind, alias))
	lines = append(lines, fmt.Sprintf("%s{%sError && <div>오류가 발생했습니다</div>}", ind, alias))
	lines = append(lines, fmt.Sprintf("%s{%s && (", ind, alias))
	lines = append(lines, fmt.Sprintf("%s  <%s%s>", ind, tag, cls))

	lines = append(lines, renderFetchJSXBody(f, alias, indent+4)...)

	lines = append(lines, fmt.Sprintf("%s  </%s>", ind, tag))
	lines = append(lines, fmt.Sprintf("%s)}", ind))

	return strings.Join(lines, "\n")
}
