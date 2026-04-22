//ff:type feature=policy type=model
//ff:what Policy — OPA Rego 정책 파싱 결과
package rego

// Policy represents parsed OPA Rego policy information.
type Policy struct {
	File       string
	Rules      []AllowRule
	Ownerships []OwnershipMapping
	ClaimsRefs []string
}
