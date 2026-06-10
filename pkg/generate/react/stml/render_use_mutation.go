//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useMutation 훅 호출 코드를 선언(data-capture/redirect/on-error) 기반으로 생성한다
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseMutation generates a useMutation hook call. The onSuccess body is
// driven entirely by the action's STML flow declarations (data-capture →
// session-store commit in bearer mode, data-redirect → navigate, neither →
// invalidateQueries). onMutate resets the action's error state on every
// (re)submission and onError feeds it — always emitted since page-flow
// Phase004 so a rejected mutation is never silent (BUG-113 (2)).
func renderUseMutation(a stmlparser.ActionBlock, fetchOps []string, bearerAuth bool, noBodyOps map[string]bool, pathParamTypes map[string]map[string]string, constraints map[string]map[string]oapiparser.FieldConstraint) string {
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	paramArgs := renderParamArgs(a.Params, a.OperationID, pathParamTypes)
	isVoid := noBodyOps[a.OperationID]

	fnParam, apiArgs := resolveMutationArgs(a.OperationID, paramArgs, isVoid, constraints)

	mutationFn := renderMutationFnExpr(fnParam, a.OperationID, apiArgs)

	captures := actionFlowCaptures(a, bearerAuth)
	onMutate := fmt.Sprintf("    onMutate: () => set%s(null),\n", toUpperFirst(errorStateVar(a)))
	onSuccess := renderOnSuccessHandler(a, captures, fetchOps)
	onError := renderOnErrorHandler(a)

	return fmt.Sprintf(`const %s = useMutation({
    mutationFn: %s,
%s%s%s  })`, mutName, mutationFn, onMutate, onSuccess, onError)
}
