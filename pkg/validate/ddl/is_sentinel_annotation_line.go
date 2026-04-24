//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what isSentinelAnnotationLine — `-- @sentinel` 주석 라인 판정 (validate 내부용)

package ddl

import (
	"strings"
)

// isSentinelAnnotationLine reports whether a trimmed line is the standalone
// `-- @sentinel` marker (allowing any amount of whitespace between `--`
// and `@sentinel`). Mirrors the parser's relaxed form so both scanners
// agree on what counts as an annotation.
func isSentinelAnnotationLine(trimmed string) bool {
	if trimmed == "-- @sentinel" {
		return true
	}
	return strings.TrimSpace(strings.TrimPrefix(trimmed, "--")) == "@sentinel"
}
