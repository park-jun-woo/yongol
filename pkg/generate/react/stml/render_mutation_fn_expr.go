//ff:func feature=stml-gen type=util control=sequence
//ff:what body only일 때 api.X 직접 참조, 그 외에는 arrow function 형태로 mutationFn 식을 반환한다
package stml

import "fmt"

// renderMutationFnExpr returns the expression assigned to the mutationFn
// property. When fnParam is empty (body only), the api function is passed
// directly (e.g. `api.Login`). Otherwise an arrow function wrapping the
// call is returned (e.g. `(data) => api.Login(data)`).
func renderMutationFnExpr(fnParam, operationID, apiArgs string) string {
	if fnParam == "" {
		return fmt.Sprintf("api.%s", operationID)
	}
	return fmt.Sprintf("%s => api.%s(%s)", fnParam, operationID, apiArgs)
}
