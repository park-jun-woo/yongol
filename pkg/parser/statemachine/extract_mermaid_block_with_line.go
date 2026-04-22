//ff:func feature=statemachine type=util control=sequence topic=states
//ff:what 마크다운에서 mermaid 코드 블록의 내용과 시작 라인 번호를 추출한다
package statemachine

import "strings"

// extractMermaidBlockWithLine extracts content from the first ```mermaid ... ``` block
// and returns the 1-based line number of the ```mermaid marker.
func extractMermaidBlockWithLine(content string) (string, int) {
	const startMarker = "```mermaid"
	const endMarker = "```"

	startIdx := strings.Index(content, startMarker)
	if startIdx < 0 {
		return "", 0
	}

	// Count newlines before startIdx to determine line number.
	lineNum := 1 + strings.Count(content[:startIdx], "\n")

	after := content[startIdx+len(startMarker):]
	endIdx := strings.Index(after, endMarker)
	if endIdx < 0 {
		return "", 0
	}
	return after[:endIdx], lineNum
}
