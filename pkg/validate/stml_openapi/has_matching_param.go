//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what hasMatchingParam — operation 에 주어진 이름의 파라미터가 있는지 확인

package stml_openapi

import "strings"

// hasMatchingParam returns true if the operation has a parameter with the
// given name (case-insensitive).
func hasMatchingParam(entry operationEntry, name string) bool {
	if entry.op == nil {
		return false
	}
	for _, p := range entry.op.Parameters {
		if p == nil || p.Value == nil {
			continue
		}
		if strings.EqualFold(p.Value.Name, name) {
			return true
		}
	}
	return false
}
