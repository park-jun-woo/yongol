//ff:func feature=stml-gen type=util control=sequence
//ff:what action에 대한 invalidation 대상 fetch operationID를 결정한다
package stml

// resolveInvalidateOps returns the fetch operationIDs to invalidate for a given
// action. If the action is inside a fetch block, only that fetch's ops are
// returned. Otherwise, all page-level fetchOps are returned.
func resolveInvalidateOps(actionOpID string, allFetchOps []string, actionFetchMap map[string][]string) []string {
	scoped, exists := actionFetchMap[actionOpID]
	if exists && scoped != nil {
		return scoped
	}
	return allFetchOps
}
