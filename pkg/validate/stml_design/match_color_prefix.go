//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-design
//ff:what matchColorPrefix — color prefix 매칭 및 토큰 기록
package stml_design

import (
	"strings"
)

// matchColorPrefix checks if the stripped class matches a color prefix and records it.
func matchColorPrefix(stripped, fullClass, file string, out *pageTokenRefs) bool {
	for _, prefix := range colorPrefixes {
		if strings.HasPrefix(stripped, prefix) {
			name := stripped[len(prefix):]
			if name == "" || isSkippable(name) {
				continue
			}
			out.Colors = append(out.Colors, tokenRef{File: file, Class: fullClass, Name: name})
			return true
		}
	}
	return false
}
