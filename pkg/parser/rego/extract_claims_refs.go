//ff:func feature=policy type=parser control=iteration dimension=1
//ff:what extractClaimsRefs — input.claims.xxx 참조를 중복 제거하여 수집
package rego

func extractClaimsRefs(content string, p *Policy) {
	seen := make(map[string]bool)
	for _, m := range reClaimsRef.FindAllStringSubmatch(content, -1) {
		if !seen[m[1]] {
			seen[m[1]] = true
			p.ClaimsRefs = append(p.ClaimsRefs, m[1])
		}
	}
}
