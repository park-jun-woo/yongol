//ff:func feature=statemachine type=util control=iteration dimension=1
//ff:what 상태명이 처음 나타나는 절대 라인 번호를 반환한다
package statemachine

import "strings"

// findStateLine returns the absolute line number where a state name first appears.
func findStateLine(lines []string, state string, mermaidStartLine int) int {
	for i, line := range lines {
		if strings.Contains(line, state) {
			return mermaidStartLine + i + 1
		}
	}
	return mermaidStartLine
}
