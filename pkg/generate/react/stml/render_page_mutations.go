//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 페이지의 모든 Action에 대한 useForm + useMutation 훅을 렌더링한다
package stml

import (
	"fmt"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageMutations(allActions []stmlparser.ActionBlock, fetchOps []string, actionFetchMap map[string][]string, constraints map[string]map[string]oapiparser.FieldConstraint, bearerAuth bool, noBodyOps map[string]bool, pathParamTypes map[string]map[string]string, sb *strings.Builder) {
	for _, a := range allActions {
		if len(a.Fields) > 0 {
			sb.WriteString(fmt.Sprintf("  %s\n", renderFormHook(a, constraints)))
		}
		// page-flow Phase004: the error state is emitted for every action
		// (data-on-error only decides the display element/position).
		errVar := errorStateVar(a)
		sb.WriteString(fmt.Sprintf("  const [%s, set%s] = useState<string | null>(null)\n", errVar, toUpperFirst(errVar)))
		targetOps := resolveInvalidateOps(a.OperationID, fetchOps, actionFetchMap, a.Invalidates)
		sb.WriteString(fmt.Sprintf("  %s\n\n", renderUseMutation(a, targetOps, bearerAuth, noBodyOps, pathParamTypes, constraints)))
	}
}
