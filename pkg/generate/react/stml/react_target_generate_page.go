//ff:func feature=stml-gen type=generator control=iteration dimension=1 topic=output
//ff:what PageSpec에서 React TSX 컴포넌트 전체 소스 코드를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func (r *ReactTarget) GeneratePage(page stmlparser.PageSpec, specsDir string, opt GenerateOptions) string {
	is := collectImports(page, specsDir)

	// Collect fetch operationIds for mutation onSuccess
	var fetchOps []string
	for _, f := range page.Fetches {
		fetchOps = collectFetchOps(f, fetchOps)
	}

	// Collect ALL actions including nested ones
	allActions := append([]stmlparser.ActionBlock{}, page.Actions...)
	allActions = append(allActions, collectAllActions(page.Children)...)
	allActions = deduplicateActions(allActions)

	// Check if any action needs a form
	needsForm := false
	for _, a := range allActions {
		if len(a.Fields) > 0 {
			needsForm = true
			break
		}
	}
	is.useForm = needsForm

	if len(allActions) > 0 {
		is.useMutation = true
		is.useQueryClient = true
	}

	var sb strings.Builder
	sb.WriteString(renderImports(is, opt))
	sb.WriteString("\n\n")

	componentName := toComponentName(page.Name)
	sb.WriteString(fmt.Sprintf("export default function %s() {\n", componentName))

	renderPageHooks(page, is, &sb)
	renderPageMutations(allActions, fetchOps, &sb)
	renderPageJSX(page, &sb)

	sb.WriteString("}\n")
	return sb.String()
}
