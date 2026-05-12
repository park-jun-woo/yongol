//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useMutation 훅 호출 코드를 생성한다
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseMutation generates a useMutation hook call.
func renderUseMutation(a stmlparser.ActionBlock, fetchOps []string, hasAuthz bool, noBodyOps map[string]bool, pathParamTypes map[string]map[string]string, constraints map[string]map[string]oapiparser.FieldConstraint) string {
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	paramArgs := renderParamArgs(a.Params, a.OperationID, pathParamTypes)
	isVoid := noBodyOps[a.OperationID]

	fnParam, apiArgs := resolveMutationArgs(a.OperationID, paramArgs, isVoid, constraints)

	// Login + authz: store tokens and navigate to '/'
	if hasAuthz && isLoginAction(a.OperationID) {
		return fmt.Sprintf(`const %s = useMutation({
    mutationFn: %s => api.%s(%s),
    onSuccess: (data) => {
      localStorage.setItem('access_token', data.access_token)
      if (data.refresh_token) {
        localStorage.setItem('refresh_token', data.refresh_token)
      }
      navigate('/')
    },
  })`, mutName, fnParam, a.OperationID, apiArgs)
	}

	invalidate := renderInvalidateExpr(fetchOps)

	return fmt.Sprintf(`const %s = useMutation({
    mutationFn: %s => api.%s(%s),
    onSuccess: () => {
      %s
    },
  })`, mutName, fnParam, a.OperationID, apiArgs, invalidate)
}
