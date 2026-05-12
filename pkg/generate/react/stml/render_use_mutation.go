//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useMutation 훅 호출 코드를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseMutation generates a useMutation hook call.
func renderUseMutation(a stmlparser.ActionBlock, fetchOps []string) string {
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	paramArgs := renderParamArgs(a.Params)

	apiArgs := "data"
	if paramArgs != "" {
		inner := strings.TrimPrefix(paramArgs, "{ ")
		inner = strings.TrimSuffix(inner, " }")
		apiArgs = "{ ...data, " + inner + " }"
	}

	// onSuccess: invalidate related queries
	invalidate := "queryClient.invalidateQueries()"
	if len(fetchOps) > 0 {
		var parts []string
		for _, op := range fetchOps {
			parts = append(parts, fmt.Sprintf("queryClient.invalidateQueries({ queryKey: ['%s'] })", op))
		}
		invalidate = strings.Join(parts, "\n      ")
	}

	return fmt.Sprintf(`const %s = useMutation({
    mutationFn: (data) => api.%s(%s),
    onSuccess: () => {
      %s
    },
  })`, mutName, a.OperationID, apiArgs, invalidate)
}
