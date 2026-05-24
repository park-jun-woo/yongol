//ff:func feature=orchestrator type=parser control=sequence
//ff:what finalizeQuerySpec — QuerySpec Body 확정 후 Params/SelectStar/SelectCols 채우기

package sqlc

import "strings"

// finalizeQuerySpec fills computed fields (Params, SelectStar, SelectCols)
// from the accumulated Body and param set. Called at both flush points
// (new `-- name:` line and end-of-file) to avoid duplication.
func finalizeQuerySpec(spec *QuerySpec, paramSet map[string]bool) {
	spec.Params = sortedKeys(paramSet)
	spec.Body = strings.TrimRight(spec.Body, " \t\r\n")
	spec.SelectStar, spec.SelectCols = extractSelectColumns(spec.Body)
}
