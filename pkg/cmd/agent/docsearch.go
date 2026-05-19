//ff:func feature=agent type=helper control=sequence
//ff:what searchDocs — diagnostic 메시지에서 rule_id/keyword 매칭으로 docs 섹션 핀포인트 추출

package agent

import (
	"strings"

	"github.com/park-jun-woo/yongol/docs"
)

// layerDocFile maps a layer to its docs filename.
func layerDocFile(l layer) string {
	switch l {
	case layerSSaC:
		return "ssac.md"
	case layerDDL, layerSQLcQuery:
		return "ddl.md"
	case layerOpenAPI:
		return "openapi.md"
	case layerRego:
		return "policy.md"
	case layerStateDiagram:
		return "states.md"
	case layerHurl:
		return "scenario.md"
	case layerManifest:
		return "manifest.md"
	case layerFuncSpec:
		return "func.md"
	default:
		return ""
	}
}

// docKeywords lists keywords to match in diagnostic messages for section lookup.
var docKeywords = []string{
	"@auth", "@call", "@state", "@get", "@post", "@put", "@delete",
	"@empty", "@exists", "@eval", "@response", "@verify-password",
	"@publish", "@subscribe",
	"BIGINT", "NOT NULL", "REFERENCES", "operationId",
}

// searchDocs returns relevant docs sections for the given layer and diagnostic messages.
// Returns empty string if no relevant section found (fallback to example only).
func searchDocs(l layer, diagMessages []string) string {
	filename := layerDocFile(l)
	if filename == "" {
		return ""
	}

	data, err := docs.FS.ReadFile(filename)
	if err != nil {
		return ""
	}

	sections := splitSections(string(data))
	if len(sections) == 0 {
		return ""
	}

	// Extract rule IDs from diagnostics
	var ruleIDs []string
	for _, msg := range diagMessages {
		rid := extractRuleID(msg)
		if rid != "" {
			ruleIDs = append(ruleIDs, rid)
		}
	}

	// Phase 1: match by rule_id
	var matched []string
	seen := map[int]bool{}
	for _, rid := range ruleIDs {
		for i, sec := range sections {
			if seen[i] {
				continue
			}
			if strings.Contains(sec, rid) {
				seen[i] = true
				matched = append(matched, sec)
			}
		}
	}

	// Phase 2: if no rule_id match, try keyword match
	if len(matched) == 0 {
		// Collect keywords present in diagnostic messages
		var activeKWs []string
		joined := strings.Join(diagMessages, "\n")
		for _, kw := range docKeywords {
			if strings.Contains(joined, kw) {
				activeKWs = append(activeKWs, kw)
			}
		}
		for _, kw := range activeKWs {
			for i, sec := range sections {
				if seen[i] {
					continue
				}
				if strings.Contains(sec, kw) {
					seen[i] = true
					matched = append(matched, sec)
				}
			}
		}
	}

	if len(matched) == 0 {
		return ""
	}

	result := strings.Join(matched, "\n\n")
	if len(result) > 2048 {
		result = result[:2048]
	}
	return result
}

// splitSections splits markdown content by "## " headings into sections.
// Each section includes its heading line.
func splitSections(content string) []string {
	lines := strings.Split(content, "\n")
	var sections []string
	var current []string

	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			if len(current) > 0 {
				sections = append(sections, strings.Join(current, "\n"))
			}
			current = []string{line}
		} else {
			if len(current) > 0 {
				current = append(current, line)
			}
			// Lines before first ## are ignored (usually just the # title)
		}
	}
	if len(current) > 0 {
		sections = append(sections, strings.Join(current, "\n"))
	}
	return sections
}
