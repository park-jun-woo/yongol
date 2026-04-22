//ff:type feature=policy type=model
//ff:what AllowRule — allow 규칙에서 추출한 액션-리소스 쌍
package rego

// AllowRule represents an extracted (action, resource) pair from an allow rule.
type AllowRule struct {
	Actions    []string
	Resource   string
	UsesOwner  bool
	UsesRole   bool
	RoleValue  string
	SourceLine int
}
