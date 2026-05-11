//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 페이지의 모든 Action에 대한 useForm + useMutation 훅을 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageMutations(allActions []stmlparser.ActionBlock, fetchOps []string, actionFetchMap map[string][]string, sb *strings.Builder) {
	for _, a := range allActions {
		if len(a.Fields) > 0 {
			sb.WriteString(fmt.Sprintf("  %s\n", renderFormHook(a)))
		}
		// Determine scoped invalidation targets for this action
		targetOps := resolveInvalidateOps(a.OperationID, fetchOps, actionFetchMap)
		sb.WriteString(fmt.Sprintf("  %s\n\n", renderUseMutation(a, targetOps)))
	}
}

// resolveInvalidateOps returns the fetch operationIDs to invalidate for a given
// action. If the action is inside a fetch block, only that fetch's ops are
// returned. Otherwise, all page-level fetchOps are returned.
func resolveInvalidateOps(actionOpID string, allFetchOps []string, actionFetchMap map[string][]string) []string {
	scoped, exists := actionFetchMap[actionOpID]
	if exists && scoped != nil {
		return scoped
	}
	// Top-level action or unknown → invalidate all page fetches
	return allFetchOps
}
