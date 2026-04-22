//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what readQueryBodyStep — scanner 한 라인 처리. 쿼리 종료 시 true 반환

package query

import (
	"strings"
)

// readQueryBodyStep handles one scanner line: initialises body.Header on the
// starting line, detects the next `-- name:` as stop, registers escape
// hatches, and appends body lines. Returns true when the caller should stop
// iterating (next query header reached).
func readQueryBodyStep(body *queryBody, line string, lineNo, startLine int, inQuery *bool) bool {
	if lineNo == startLine {
		body.Header = line
		*inQuery = true
		return false
	}
	if !*inQuery {
		return false
	}
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "-- name:") {
		return true
	}
	// Detect escape hatches in adjacent comment lines. Accept both the
	// `+<name>` form (safe inside a SQL comment) and the `@<name>` form
	// (hyphenated, so sqlc named-param regex stops before the hyphen
	// and does not misinterpret the marker — we still register both to
	// keep backward-compat with docs that use @).
	if strings.HasPrefix(trimmed, "--") {
		registerQueryBodyEscapes(body, trimmed)
	}
	body.Lines = append(body.Lines, line)
	return false
}
