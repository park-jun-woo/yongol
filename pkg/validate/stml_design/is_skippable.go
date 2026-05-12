//ff:func feature=validate type=util control=selection topic=stml-design
//ff:what isSkippable — Tailwind 내장 값/숫자/팔레트 패턴 등 커스텀 토큰 아닌 값 판별
package stml_design

import (
	"strings"
)

// isSkippable returns true for values that cannot be custom DESIGN.md tokens:
// numeric values, Tailwind builtins, palette patterns (gray-500), and arbitrary values.
func isSkippable(s string) bool {
	if s == "" {
		return false
	}
	// Arbitrary value [...]
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		return true
	}
	// Standard Tailwind palette pattern (e.g. "gray-500", "red-200")
	if tailwindPaletteRe.MatchString(s) {
		return true
	}
	// Common Tailwind built-in keywords
	switch s {
	case "full", "none", "auto", "px", "white", "black", "inherit",
		"current", "transparent", "thin", "extralight", "light",
		"normal", "medium", "semibold", "bold", "extrabold",
		"screen", "fit", "min", "max", "prose",
		"sans", "serif", "mono":
		return true
	}
	// Pure numeric (integer, decimal, or fraction: "4", "0.5", "1/2")
	return isPureNumeric(s)
}
