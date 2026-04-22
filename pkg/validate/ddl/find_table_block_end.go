//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what findTableBlockEnd — locate the closing line index of a CREATE TABLE block

package ddl

import "strings"

// findTableBlockEnd scans lines starting at startIdx+1 for the block
// terminator (a line starting with ")") and returns the end index. When no
// terminator is found, returns the last valid index so the caller can still
// slice a block body.
func findTableBlockEnd(lines []string, startIdx int) int {
	j := startIdx + 1
	for j < len(lines) {
		trimmed := strings.TrimSpace(lines[j])
		if strings.HasPrefix(trimmed, ")") {
			return j
		}
		j++
	}
	if j >= len(lines) {
		return len(lines) - 1
	}
	return j
}
