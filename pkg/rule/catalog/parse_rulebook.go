//ff:func feature=rule type=parser control=iteration dimension=2 topic=catalog
//ff:what Parse — rulebook.md 의 H2 섹션 + `| Rule ID | Level | Description | Source |` 테이블을 RuleMeta 슬라이스로
package catalog

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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
// leading/trailing backticks (since Source is typically wrapped in `` ` ``).
func Parse(r io.Reader) ([]RuleMeta, error) {
	scanner := bufio.NewScanner(r)
	// rulebook rows are short; default buffer size is plenty but raise the
	// cap defensively in case a single description balloons.
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var (
		rules        []RuleMeta
		sectionTitle string
		sectionSkip  bool // true while inside ## Deprecated
		inTable      bool // past header + separator rows
		sawHeader    bool // saw "| Rule ID | Level | ..." row
	)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Section transitions.
		if strings.HasPrefix(trimmed, "## ") {
			sectionTitle = strings.TrimSpace(strings.TrimPrefix(trimmed, "##"))
			sectionSkip = strings.EqualFold(sectionTitle, "Deprecated")
			inTable = false
			sawHeader = false
			continue
		}
		// H3 / lesser headings don't reset section state but do terminate
		// a rule table (a new subheading means the table ended).
		if strings.HasPrefix(trimmed, "### ") || strings.HasPrefix(trimmed, "#### ") {
			inTable = false
			sawHeader = false
			continue
		}

		if sectionSkip {
			continue
		}

		// Blank line ends a table context cheaply.
		if trimmed == "" {
			inTable = false
			sawHeader = false
			continue
		}

		if !strings.HasPrefix(trimmed, "|") {
			// Non-table line inside the section body — harmless prose.
			inTable = false
			sawHeader = false
			continue
		}

		// Inside a table row.
		if !sawHeader {
			// Looking for "| Rule ID | Level | Description | Source |" header.
			if isRuleTableHeader(trimmed) {
				sawHeader = true
			}
			continue
		}
		if !inTable {
			// Header seen; next `|---|...|` separator row switches to data mode.
			if isTableSeparator(trimmed) {
				inTable = true
			} else {
				// Unexpected row between header and separator — bail on this table.
				sawHeader = false
			}
			continue
		}

		// Data row.
		cells := splitRow(trimmed)
		if len(cells) < 4 {
			continue
		}
		id := strings.TrimSpace(cells[0])
		level := strings.ToUpper(strings.TrimSpace(cells[1]))
		desc := strings.TrimSpace(cells[2])
		source := strings.Trim(strings.TrimSpace(cells[3]), "`")

		if id == "" || (level != "ERROR" && level != "WARNING") {
			continue
		}

		rules = append(rules, RuleMeta{
			ID:            id,
			Level:         level,
			Description:   desc,
			Source:        source,
			SectionTitle:  sectionTitle,
			SectionAnchor: sectionAnchor(sectionTitle),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("rulebook scan: %w", err)
	}
	return rules, nil
}
