//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what splitRolesAttr — data-roles 값("a, b")을 트림된 role 목록으로 분해 (빈 항목 제외)

package stml

import "strings"

// splitRolesAttr splits a data-roles attribute value into its role names:
// comma-separated, whitespace-trimmed, empty entries dropped. An empty or
// absent attribute yields nil — the node renders unconditionally.
func splitRolesAttr(raw string) []string {
	if raw == "" {
		return nil
	}
	var roles []string
	for _, seg := range strings.Split(raw, ",") {
		if seg = strings.TrimSpace(seg); seg != "" {
			roles = append(roles, seg)
		}
	}
	return roles
}
