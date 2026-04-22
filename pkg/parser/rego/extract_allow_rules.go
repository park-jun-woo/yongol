//ff:func feature=policy type=parser control=iteration dimension=1
//ff:what extractAllowRules — splits allow blocks from Rego source and extracts AllowRule entries
package rego

import "strings"

func extractAllowRules(content string, p *Policy) {
	normalized := strings.ReplaceAll(content, "\nallow {", "\nallow if {")
	// Track cumulative newline counts to compute 1-based line numbers relative to the original content.
	// The "\nallow {" → "\nallow if {" substitution only adds tokens within the same line, so line numbers are preserved.
	const token = "allow if {"
	offset := 0
	for {
		idx := strings.Index(normalized[offset:], token)
		if idx < 0 {
			break
		}
		headerStart := offset + idx
		bodyStart := headerStart + len(token)
		rel := findClosingBrace(normalized[bodyStart:])
		if rel < 0 {
			offset = bodyStart
			continue
		}
		endIdx := bodyStart + rel
		block := normalized[bodyStart:endIdx]
		if rule, ok := processAllowBlock(block); ok {
			rule.SourceLine = lineOfOffset(normalized, headerStart)
			p.Rules = append(p.Rules, rule)
		}
		offset = endIdx + 1
	}
}

// lineOfOffset returns the 1-based line number at byte offset off in string s.
// Returns 0 when off is out of range.
func lineOfOffset(s string, off int) int {
	if off < 0 || off > len(s) {
		return 0
	}
	return 1 + strings.Count(s[:off], "\n")
}
