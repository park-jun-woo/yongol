//ff:func feature=projectconfig type=parser control=iteration dimension=1
//ff:what collectClaimLines — auth.claims mapping 에서 각 claim 키의 줄 번호를 수집
package manifest

import (
	"gopkg.in/yaml.v3"
)

// collectClaimLines walks auth.claims (MappingNode) and records each claim
// key's 1-based line number into the caller-supplied map.
func collectClaimLines(auth *yaml.Node, out map[string]int) {
	claims := mappingValue(auth, "claims")
	if claims == nil || claims.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(claims.Content); i += 2 {
		k := claims.Content[i]
		out[k.Value] = k.Line
	}
}
