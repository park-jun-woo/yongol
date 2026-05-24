//ff:func feature=orchestrator type=parser control=sequence
//ff:what processSQLCScanLine — sqlc 파일 한 줄을 해석해 QuerySpec 전이 및 named param/Body 수집을 처리
package sqlc

// processSQLCScanLine handles a single line of a sqlc source file during the
// streaming scan. It may open a new QuerySpec (flushing the previous one into
// specs) or accumulate the line into the current spec's Body and named params.
// It returns the updated (current, paramSet).
//
// Body accumulation: every non-`-- name:` line that follows a `-- name:` marker
// is appended to current.Body (joined with "\n"). The marker line itself is
// not included. On flush, trailing whitespace is trimmed.
func processSQLCScanLine(
	line, model, path string,
	lineNo int,
	current *QuerySpec,
	paramSet map[string]bool,
	specs *[]QuerySpec,
) (*QuerySpec, map[string]bool) {
	if spec, ok := parseSQLCLine(line, model, path, lineNo); ok {
		// Flush previous query
		if current != nil {
			finalizeQuerySpec(current, paramSet)
			*specs = append(*specs, *current)
		}
		newSpec := spec
		return &newSpec, make(map[string]bool)
	}
	if current != nil {
		appendBodyLine(current, line)
		collectNamedParams(line, paramSet)
	}
	return current, paramSet
}
