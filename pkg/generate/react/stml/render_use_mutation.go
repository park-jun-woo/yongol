//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useMutation 훅 호출 코드를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseMutation generates a useMutation hook call.
func renderUseMutation(a stmlparser.ActionBlock, fetchOps []string, hasAuthz bool, noBodyOps map[string]bool, pathParamTypes map[string]map[string]string) string {
	mutName := toLowerFirst(a.OperationID) + "Mutation"
	paramArgs := renderParamArgs(a.Params, a.OperationID, pathParamTypes)
	isVoid := noBodyOps[a.OperationID]

	var fnParam, apiArgs string
	if isVoid {
		fnParam = "()"
		if paramArgs != "" {
			apiArgs = paramArgs
		} else {
			apiArgs = ""
		}
	} else {
		fnParam = "(data)"
		apiArgs = "data"
		if paramArgs != "" {
			inner := strings.TrimPrefix(paramArgs, "{ ")
			inner = strings.TrimSuffix(inner, " }")
			apiArgs = "{ ...data, " + inner + " }"
		}
	}

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
    mutationFn: %s => api.%s(%s),
    onSuccess: () => {
      %s
    },
  })`, mutName, fnParam, a.OperationID, apiArgs, invalidate)
}
