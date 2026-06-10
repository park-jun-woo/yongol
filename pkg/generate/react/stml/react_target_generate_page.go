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
	// Pre-populate row-action mutate arguments (item.* sources, Phase006)
	populateRowActionArgs(&page, opt.ResponseArrayItemTypes, opt.PathParamTypes)
	// Pre-resolve data-link target route patterns (Phase007)
	populateLinkTargets(&page, opt.RoutePatterns)

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
	// Pre-resolve page-name data-redirect target route patterns (Phase008)
	for i := range allActions {
		setRedirectPattern(&allActions[i], opt.RoutePatterns)
	}

	is.useForm = anyActionHasFields(allActions)
	is.useZod = is.useForm && anyActionHasZodConstraints(allActions, opt.RequestConstraints)

	if len(allActions) > 0 {
		is.useMutation = true
		is.useButton = true
		// page-flow Phase004: every action keeps an error-message state
		// (default onError emission), so useState is needed regardless of
		// data-on-error declarations.
		is.useState = true
	}
	if anyActionHasInputFields(allActions) {
		is.useInput = true
	}

	// STML flow declarations drive imports: data-redirect → useNavigate,
	// data-capture (bearer) → session store. queryClient is needed only
	// when at least one action keeps the default invalidateQueries() path.
	needsInvalidate := false
	for _, a := range allActions {
		if a.Redirect != "" {
			is.useNavigate = true
		}
		// route.<Name> redirect params read the useParams() variable
		// (page-flow Phase008, same need as data-link route.* sources).
		if redirectUsesRouteParams(a) {
			is.useParams = true
		}
		if len(actionFlowCaptures(a, opt.BearerAuth)) > 0 {
			is.useAuthStore = true
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
