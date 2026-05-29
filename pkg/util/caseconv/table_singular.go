//ff:func feature=util type=util control=selection topic=string-convert
//ff:what TableSingular — 복수형 lower-snake 테이블명 → 단수형 (생성기·검증기 공유)

package caseconv

import "strings"

// TableSingular desingularises a lower-snake table name to its singular lower
// form. Shared by code generators (gogin/ssac, ir) and the XSD-55 validator so
// model↔table matching stays identical across layers. Input is assumed to be
// lower-snake; the rules mirror inflection.Singular on common fixtures
// (users / organizations / companies / addresses / boxes).
func TableSingular(name string) string {
	switch {
	case strings.HasSuffix(name, "ies"):
		return name[:len(name)-3] + "y"
	case strings.HasSuffix(name, "sses"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "xes"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "s") && !strings.HasSuffix(name, "ss"):
		return name[:len(name)-1]
	default:
		return name
	}
}
