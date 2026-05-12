//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what matchRoundedPrefix — rounded prefix 매칭 및 토큰 기록
package stml_design

import (
	"strings"
)

// matchRoundedPrefix checks if the stripped class matches the rounded prefix.
func matchRoundedPrefix(stripped, fullClass, file string, out *pageTokenRefs) bool {
	if !strings.HasPrefix(stripped, "rounded-") {
		return false
	}
	name := stripped[len("rounded-"):]
	if name == "" || isSkippable(name) {
		return false
	}
	out.Rounded = append(out.Rounded, tokenRef{File: file, Class: fullClass, Name: name})
	return true
}
