//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what TypesCompatible — 두 Go 타입 이름 문자열이 호환되는지 판정 (공유 헬퍼)

package ssac_func

import "strings"

// TypesCompatible reports whether actual can be assigned to expected,
// under Go's runtime/generated-code semantics as modelled by yongol.
//
// Rules:
//   - exact string equality
//   - pointer prefix `*` stripped before comparison
//   - int family interchangeable (int, int8~64, uint*, byte, rune)
//   - nil is compatible with any side (pointer/interface/slice/map is
//     structurally valid target; we accept optimistically)
//   - object ↔ primitive are NOT compatible (falls through to false)
//
// Used by XFS-44 (@call input type) and XOS-67 (@response field type).
func TypesCompatible(actual, expected string) bool {
	a := strings.TrimPrefix(actual, "*")
	e := strings.TrimPrefix(expected, "*")
	if a == e {
		return true
	}
	if isIntType(a) && isIntType(e) {
		return true
	}
	if a == "nil" || e == "nil" {
		return true
	}
	return false
}
