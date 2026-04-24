//ff:func feature=validate type=util control=iteration dimension=2 topic=ddl-structural
//ff:what scanInsertsWithAnnotations — validate 로컬 INSERT 스캐너 (parser 스캐너 미러)

package ddl

import (
	"strings"
)

// scanInsertsWithAnnotations is the validate-local copy of the parser's
// sentinel scanner. It walks content line by line, locates top-level
// INSERT statements, and returns each paired with whether a `-- @sentinel`
// annotation preceded it. The body is collected through an unquoted `;`.
//
// The parser owns the authoritative implementation (pkg/parser/ddl). This
// copy exists only to keep pkg/validate/ddl free of an import cycle on
// the parser package; both must agree on what "top-level INSERT" means.
func scanInsertsWithAnnotations(content string) []insertScan {
	lines := strings.Split(content, "\n")
	var results []insertScan
	annotated := false
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			i++
			continue
		}
		if isSentinelAnnotationLine(trimmed) {
			annotated = true
			i++
			continue
		}
		if m := insertIntoLineRe.FindStringSubmatch(lines[i]); m != nil {
			r, next := collectValidateInsertScan(lines, i, m[1], annotated)
			results = append(results, r)
			annotated = false
			i = next
			continue
		}
		annotated = false
		i++
	}
	return results
}
