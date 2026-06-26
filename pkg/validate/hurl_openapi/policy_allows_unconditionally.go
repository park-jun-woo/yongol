//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what policyAllowsUnconditionally — 단일 Policy 내 action+resource 매칭 규칙의 조건부 여부 판정

package hurl_openapi

import (
	"slices"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func policyAllowsUnconditionally(pol rego.Policy, action, resource string) (found, conditional bool) {
	for _, rule := range pol.Rules {
		if rule.Resource != resource || !slices.Contains(rule.Actions, action) {
			continue
		}
		if rule.UsesOwner || rule.UsesRole {
			return true, true
		}
		found = true
	}
	return
}
