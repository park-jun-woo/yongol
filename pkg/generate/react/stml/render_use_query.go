//ff:func feature=stml-gen type=generator control=sequence
//ff:what FetchBlock에 대한 useQuery 훅 호출 코드를 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseQuery generates a useQuery hook call.
func renderUseQuery(f stmlparser.FetchBlock, pathParamTypes map[string]map[string]string) string {
	alias := toLowerFirst(f.OperationID) + "Data"
	paramValues := renderParamValues(f.Params)
	paramArgs := renderParamArgs(f.Params, f.OperationID, pathParamTypes)

	// queryKey parts
	queryKey := fmt.Sprintf("'%s'", f.OperationID)
	if paramValues != "" {
		queryKey += ", " + paramValues
	}

	// API call args
	apiCall := fmt.Sprintf("api.%s(%s)", f.OperationID, paramArgs)

	// Optional route params can be absent — gate the query so it does not fire
	// with NaN / undefined before the segment is present (BUG-136). Required
	// params always exist in the matched route, so they emit no guard
	// (no regression to existing snapshots).
	enabled := renderEnabledGuard(f, pathParamTypes)

	return fmt.Sprintf(`const { data: %s, isLoading: %sLoading, error: %sError } = useQuery({
    queryKey: [%s],
    queryFn: () => %s,%s
  })`, alias, alias, alias, queryKey, apiCall, enabled)
}
