//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what isAllowAll — action+resource 의 OPA 규칙이 무조건 allow 인지 판정

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func isAllowAll(action, resource string, policies []rego.Policy) bool {
	found := false
	for _, pol := range policies {
		f, conditional := policyAllowsUnconditionally(pol, action, resource)
		if conditional {
			return false
		}
		if f {
			found = true
		}
	}
	return found
}
