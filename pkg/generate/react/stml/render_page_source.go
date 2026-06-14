//ff:func feature=stml-gen type=generator control=sequence topic=output
//ff:what 해석된 import/title/crumb 정보로 페이지 컴포넌트의 import·함수 선언·hooks·mutations·JSX 전체 소스 문자열을 조립한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderPageSource assembles the final TSX source from the import set and the
// resolved per-page flags. It is the pure rendering tail of GeneratePage —
// extracted verbatim so that function stays within the per-func line limit;
// the call order (imports → component header → title effect → hooks → crumb
// effect → mutations → JSX) is preserved.
func renderPageSource(page stmlparser.PageSpec, is importSet, allActions []stmlparser.ActionBlock, fetchOps []string, actionFetchMap map[string][]string, docTitle, crumbField string, opt GenerateOptions) string {
	var sb strings.Builder
	sb.WriteString(renderImports(is, opt))
	sb.WriteString("\n\n")

	componentName := toComponentName(page.Name)
	sb.WriteString(fmt.Sprintf("export default function %s() {\n", componentName))

	if docTitle != "" {
		sb.WriteString(renderTitleEffect(docTitle))
	}
	renderPageHooks(page, is, opt.PathParamTypes, &sb)
	if crumbField != "" {
		sb.WriteString(renderCrumbLabelEffect(crumbField, toLowerFirst(page.Fetches[0].OperationID)+"Data", opt.CrumbTitleSuffix))
	}
	renderPageMutations(allActions, fetchOps, actionFetchMap, opt.RequestConstraints, opt.BearerAuth, opt.NoBodyOps, opt.PathParamTypes, opt.ErrorDisplayField, opt.ResponseBindTypes, &sb)
	renderPageJSX(page, &sb, opt.NoBodyOps, bindCtx{all: opt.ResponseBindTypes})

	sb.WriteString("}\n")
	return sb.String()
}
