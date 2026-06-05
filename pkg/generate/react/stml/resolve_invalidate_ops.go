//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what action에 대한 invalidation 대상 fetch operationID를 결정한다 (자동 추론 + 명시 병합)
package stml

// resolveInvalidateOps returns the fetch operationIDs to invalidate for a given
// action. The auto-inferred set is the action's scoped fetch ops (when inside a
// fetch block) or all page-level fetchOps otherwise. explicitInvalidates
// (data-invalidates) is merged in via union with duplicate removal, preserving
// the auto-inferred order first. When explicitInvalidates is empty the output
// equals the auto-inferred set unchanged.
func resolveInvalidateOps(actionOpID string, allFetchOps []string, actionFetchMap map[string][]string, explicitInvalidates []string) []string {
	auto := allFetchOps
	if scoped, exists := actionFetchMap[actionOpID]; exists && scoped != nil {
		auto = scoped
	}
	if len(explicitInvalidates) == 0 {
		return auto
	}
	seen := make(map[string]bool, len(auto))
	merged := make([]string, 0, len(auto)+len(explicitInvalidates))
	for _, op := range auto {
		if seen[op] {
			continue
		}
		seen[op] = true
		merged = append(merged, op)
	}
	for _, op := range explicitInvalidates {
		if seen[op] {
			continue
		}
		seen[op] = true
		merged = append(merged, op)
	}
	return merged
}
