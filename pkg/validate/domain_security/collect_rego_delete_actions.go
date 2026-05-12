//ff:func feature=validate type=util control=iteration dimension=2 topic=domain-security
//ff:what collectRegoDeleteActions — Rego 정책에서 "delete" action 리소스 수집
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectRegoDeleteActions gathers resources that have "delete" actions in Rego policies.
func collectRegoDeleteActions(fs *yongol.Fullstack) map[string]struct{} {
	result := make(map[string]struct{})
	for _, policy := range fs.ParsedPolicies {
		for _, rule := range policy.Rules {
			if hasDeleteAction(rule.Actions) {
				result[rule.Resource] = struct{}{}
			}
		}
	}
	return result
}
