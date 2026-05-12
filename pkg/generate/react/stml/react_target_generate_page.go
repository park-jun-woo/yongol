//ff:func feature=stml-gen type=generator control=sequence topic=output
//ff:what PageSpec에서 React TSX 컴포넌트 전체 소스 코드를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func (r *ReactTarget) GeneratePage(page stmlparser.PageSpec, specsDir string, opt GenerateOptions) string {
	// Pre-populate EachBlock.KeyField from response schema
	populateEachKeyFields(&page, opt.ResponseArrayItemFields)

	is := collectImports(page, specsDir)

	// Collect fetch operationIds for mutation onSuccess (all page-level fetches)
	var fetchOps []string
	for _, f := range page.Fetches {
		fetchOps = collectFetchOps(f, fetchOps)
	}

	// Build per-action fetch map for scoped invalidation
	actionFetchMap := buildActionFetchMap(page)

	// Collect ALL actions including nested ones
	allActions := append([]stmlparser.ActionBlock{}, page.Actions...)
	allActions = append(allActions, collectAllActions(page.Children)...)
	allActions = deduplicateActions(allActions)

	is.useForm = anyActionHasFields(allActions)
	is.useZod = is.useForm && anyActionHasZodConstraints(allActions, opt.RequestConstraints)

	if len(allActions) > 0 {
		is.useMutation = true
		is.useQueryClient = true
		is.useButton = true
	}
	if anyActionHasInputFields(allActions) {
		is.useInput = true
	}

	// Login + authz: navigate on success instead of invalidate
	if opt.HasAuthz && hasLoginAction(allActions) {
		is.useNavigate = true
		// Drop queryClient only when every action is Login (no invalidation needed)
		if allLoginActions(allActions) {
			is.useQueryClient = false
		}
	}

	var sb strings.Builder
	sb.WriteString(renderImports(is, opt))
	sb.WriteString("\n\n")

	componentName := toComponentName(page.Name)
	sb.WriteString(fmt.Sprintf("export default function %s() {\n", componentName))

	renderPageHooks(page, is, &sb)
	renderPageMutations(allActions, fetchOps, actionFetchMap, opt.RequestConstraints, opt.HasAuthz, &sb)
	renderPageJSX(page, &sb, opt.NoBodyOps)

	sb.WriteString("}\n")
	return sb.String()
}
