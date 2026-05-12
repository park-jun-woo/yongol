//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what hasRegoDeleteRule — operationId가 Rego delete 룰에 의해 커버되는지 확인
package domain_security

// hasRegoDeleteRule checks if the operationId is covered by any Rego delete rule.
// It checks both exact resource match and plural/singular variations.
func hasRegoDeleteRule(operationID string, regoResources map[string]struct{}) bool {
	if len(regoResources) == 0 {
		return false
	}
	// Check if any Rego resource name is a substring of the operationId (case-insensitive match).
	// E.g., operationId "DeleteWorkflow" should match resource "workflow".
	opLower := toLower(operationID)
	for resource := range regoResources {
		if contains(opLower, toLower(resource)) {
			return true
		}
	}
	return false
}
