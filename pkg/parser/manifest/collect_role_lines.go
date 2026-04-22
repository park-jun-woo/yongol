//ff:func feature=projectconfig type=parser control=iteration dimension=1
//ff:what collectRoleLines — auth.roles sequence 에서 각 role 값의 줄 번호를 수집
package manifest

import (
	"gopkg.in/yaml.v3"
)

// collectRoleLines walks auth.roles (SequenceNode) and records each role
// value's 1-based line number into the caller-supplied map.
func collectRoleLines(auth *yaml.Node, out map[string]int) {
	roles := mappingValue(auth, "roles")
	if roles == nil || roles.Kind != yaml.SequenceNode {
		return
	}
	for _, item := range roles.Content {
		out[item.Value] = item.Line
	}
}
