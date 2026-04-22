//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildOperationRoleMap — fs.ParsedPolicies에서 operation → role 매핑 추출
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// buildOperationRoleMap walks parsed OPA policies and returns action → role.
// Only rules with UsesRole && RoleValue != "" contribute.
func buildOperationRoleMap(policies []rego.Policy) map[string]string {
	roleMap := map[string]string{}
	for _, p := range policies {
		for _, r := range p.Rules {
			if !r.UsesRole || r.RoleValue == "" {
				continue
			}
			for _, action := range r.Actions {
				roleMap[action] = r.RoleValue
			}
		}
	}
	return roleMap
}
