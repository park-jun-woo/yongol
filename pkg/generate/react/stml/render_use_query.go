//ff:func feature=stml-gen type=generator control=sequence
//ff:what FetchBlock에 대한 useQuery 훅 호출 코드를 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseQuery generates a useQuery hook call.
func renderUseQuery(f stmlparser.FetchBlock) string {
	alias := toLowerFirst(f.OperationID) + "Data"
	paramValues := renderParamValues(f.Params)
	paramArgs := renderParamArgs(f.Params)

	// queryKey parts
	queryKey := fmt.Sprintf("'%s'", f.OperationID)
	if paramValues != "" {
		queryKey += ", " + paramValues
	}

	// API call args
	apiCall := fmt.Sprintf("api.%s(%s)", f.OperationID, paramArgs)

	return fmt.Sprintf(`const { data: %s, isLoading: %sLoading, error: %sError } = useQuery({
    queryKey: [%s],
    queryFn: () => %s,
  })`, alias, alias, alias, queryKey, apiCall)
}
