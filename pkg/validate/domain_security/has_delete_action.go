//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what hasDeleteAction — AllowRule의 actions에 "delete"가 있는지 확인
package domain_security

// hasDeleteAction checks if a list of actions contains "delete".
func hasDeleteAction(actions []string) bool {
	for _, action := range actions {
		if action == "delete" {
			return true
		}
	}
	return false
}
