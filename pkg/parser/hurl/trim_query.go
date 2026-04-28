//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what trimQuery — URL path 끝의 ?query fragment 제거

package hurl

import "strings"

// trimQuery strips a trailing `?query` fragment from a URL path.
func trimQuery(p string) string {
	if idx := strings.Index(p, "?"); idx >= 0 {
		return p[:idx]
	}
	return p
}
