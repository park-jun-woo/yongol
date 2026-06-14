//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what action의 invalidate 대상 GET을 결정하고, delete는 자기 GET을 removeQueries 대상으로 분리한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// resolveInvalidateOps returns the fetch operationIDs to refetch (invalidate)
// and, for a delete action, the deleted item's own GET to drop (remove) on
// success. The auto-inferred set is the action's scoped fetch ops (when
// inside a fetch block) or all page-level fetchOps otherwise; explicit
// data-invalidates is merged in via union with duplicate removal, preserving
// the auto-inferred order first. For a delete action the self GET (same
// path-param signature, BUG-132 132-2) is split out of invalidate into
// remove — invalidating it would refetch a 404. Non-delete actions return a
// nil remove set.
func resolveInvalidateOps(a stmlparser.ActionBlock, allFetchOps []string, actionFetchMap map[string][]string, pathParamTypes map[string]map[string]string) (invalidate, remove []string) {
	auto := allFetchOps
	if scoped, exists := actionFetchMap[a.OperationID]; exists && scoped != nil {
		auto = scoped
	}
	seen := make(map[string]bool, len(auto)+len(a.Invalidates))
	merged := make([]string, 0, len(auto)+len(a.Invalidates))
	for _, op := range auto {
		if seen[op] {
			continue
		}
		seen[op] = true
		merged = append(merged, op)
	}
	for _, op := range a.Invalidates {
		if seen[op] {
			continue
		}
		seen[op] = true
		merged = append(merged, op)
	}

	if !isDeleteOperation(a.OperationID) {
		return merged, nil
	}
	selfGets := selfGetOps(a.OperationID, merged, pathParamTypes)
	if len(selfGets) == 0 {
		return merged, nil
	}
	removeSet := make(map[string]bool, len(selfGets))
	for _, op := range selfGets {
		removeSet[op] = true
	}
	invalidate = make([]string, 0, len(merged))
	for _, op := range merged {
		if removeSet[op] {
			continue
		}
		invalidate = append(invalidate, op)
	}
	return invalidate, selfGets
}
