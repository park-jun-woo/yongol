//ff:func feature=statemachine type=util control=sequence topic=states
//ff:what 마크다운에서 mermaid 코드 블록의 내용을 추출한다
package statemachine

import "strings"

// extractMermaidBlock extracts content from the first ```mermaid ... ``` block.
func extractMermaidBlock(content string) string {
	const startMarker = "```mermaid"
	const endMarker = "```"

	startIdx := strings.Index(content, startMarker)
	if startIdx < 0 {
		return ""
	}
	after := content[startIdx+len(startMarker):]
	endIdx := strings.Index(after, endMarker)
	if endIdx < 0 {
		return ""
	}
	return after[:endIdx]
}
