//ff:func feature=validate type=util control=selection topic=func-check
//ff:what isIntType — Go int 계열 타입명(t)인지 판정

package ssac_func

// isIntType reports whether t is any of Go's int family names.
func isIntType(t string) bool {
	switch t {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "byte", "rune":
		return true
	}
	return false
}
