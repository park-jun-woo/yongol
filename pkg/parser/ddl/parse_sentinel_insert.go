//ff:func feature=manifest type=parser control=iteration dimension=2
//ff:what parseSentinelInserts — 파일 본문에서 top-level INSERT 블록을 찾아 @sentinel 어노테이션과 묶어 반환
package ddl

import (
	"strings"
)

// parseSentinelInserts walks the raw file content line by line, finds
// top-level INSERT statements, and returns each one paired with whether
// the line immediately above it (skipping blank lines) carried the
// `-- @sentinel` annotation. The INSERT body is collected verbatim
// through its terminating `;`. Strings in single quotes are respected so
// a `;` inside a literal does not end the statement prematurely.
func parseSentinelInserts(content string) []sentinelScanResult {
	lines := strings.Split(content, "\n")
	var results []sentinelScanResult
	annotated := false
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			// blank lines do not reset the annotation (allows one blank between)
			i++
			continue
		}
		if isSentinelAnnotation(trimmed) {
			annotated = true
			i++
			continue
		}
		if m := insertIntoRe.FindStringSubmatch(lines[i]); m != nil {
			r, next := collectSentinelInsert(lines, i, m[1], annotated)
			results = append(results, r)
			annotated = false
			i = next
			continue
		}
		// any other non-annotation line invalidates a pending annotation
		annotated = false
		i++
	}
	return results
}
