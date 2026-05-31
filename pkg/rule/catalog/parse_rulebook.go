//ff:func feature=rule type=parser control=iteration dimension=2 topic=catalog
//ff:what Parse — rulebook.md 의 H2 섹션 + `| Rule ID | Level | Description | Source |` 테이블을 RuleMeta 슬라이스로
package catalog

import (
	"bufio"
	"fmt"
	"io"
)

// Parse reads rulebook.md content from r and returns the ordered list of
// RuleMeta rows. Rows inside the `## Deprecated` section are skipped.
//
// Expected format:
//
//	## A. SSaC Internal
//	...
//	| Rule ID | Level | Description | Source |
//	|---|---|---|---|
//	| S-27 | ERROR | Variable must be declared before use | `pkg/validate/ssac/s_27_var_declared.go` |
//
// H2 headings prefixed with `##` open a new section. The first table after
// a section header whose header row starts with `| Rule ID` is consumed.
// Subsequent tables inside the same section (currently none) are ignored.
//
// Description and Source fields are trimmed of surrounding whitespace and
// leading/trailing backticks (since Source is typically wrapped in “ ` “).
func Parse(r io.Reader) ([]RuleMeta, error) {
	scanner := bufio.NewScanner(r)
	// rulebook rows are short; default buffer size is plenty but raise the
	// cap defensively in case a single description balloons.
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var st rulebookParseState
	for scanner.Scan() {
		st.feedLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("rulebook scan: %w", err)
	}
	return st.rules, nil
}
