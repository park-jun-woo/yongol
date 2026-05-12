//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-design
//ff:what matchSpacingPrefix — spacing prefix 매칭 및 토큰 기록
package stml_design

import (
	"strings"
)

// matchSpacingPrefix checks if the stripped class matches a spacing prefix and records it.
func matchSpacingPrefix(stripped, fullClass, file string, out *pageTokenRefs) bool {
	for _, prefix := range spacingPrefixes {
		if strings.HasPrefix(stripped, prefix) {
			name := stripped[len(prefix):]
			if name == "" || isSkippable(name) {
				continue
			}
			out.Spacing = append(out.Spacing, tokenRef{File: file, Class: fullClass, Name: name})
			return true
		}
	}
	return false
}
