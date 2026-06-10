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
		is.useButton = true
	}
	if anyActionHasInputFields(allActions) {
		is.useInput = true
	}

	// STML flow declarations drive imports: data-redirect → useNavigate,
	// data-capture (bearer) → session store, data-on-error → useState.
	// queryClient is needed only when at least one action keeps the
	// default invalidateQueries() path.
	needsInvalidate := false
	for _, a := range allActions {
		if a.Redirect != "" {
			is.useNavigate = true
		}
		if len(actionFlowCaptures(a, opt.BearerAuth)) > 0 {
			is.useAuthStore = true
		}
		if a.OnErrorNode {
			is.useState = true
		}
		if !actionHasFlowSuccess(a, opt.BearerAuth) {
			needsInvalidate = true
		}
	}
	is.useQueryClient = len(allActions) > 0 && needsInvalidate

	var sb strings.Builder
	sb.WriteString(renderImports(is, opt))
	sb.WriteString("\n\n")

	componentName := toComponentName(page.Name)
	sb.WriteString(fmt.Sprintf("export default function %s() {\n", componentName))

	renderPageHooks(page, is, opt.PathParamTypes, &sb)
	renderPageMutations(allActions, fetchOps, actionFetchMap, opt.RequestConstraints, opt.BearerAuth, opt.NoBodyOps, opt.PathParamTypes, &sb)
	renderPageJSX(page, &sb, opt.NoBodyOps)

	sb.WriteString("}\n")
	return sb.String()
}
