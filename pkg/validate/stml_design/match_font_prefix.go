//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what matchFontPrefix — font prefix 매칭 및 토큰 기록
package stml_design

import (
	"strings"
)

// matchFontPrefix checks if the stripped class matches the font prefix.
func matchFontPrefix(stripped, fullClass, file string, out *pageTokenRefs) {
	if !strings.HasPrefix(stripped, "font-") {
		return
	}
	name := stripped[len("font-"):]
	if name == "" || isSkippable(name) {
		return
	}
	out.Fonts = append(out.Fonts, tokenRef{File: file, Class: fullClass, Name: name})
}
