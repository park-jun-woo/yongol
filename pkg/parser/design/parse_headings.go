//ff:func feature=frontend type=parser control=iteration dimension=1
//ff:what Markdown body 에서 ## 레벨 헤딩 텍스트를 순서대로 추출
package design

import (
	"bufio"
	"bytes"
	"strings"
)

// parseHeadings extracts all level-2 headings (## ...) from a Markdown body.
func parseHeadings(body []byte) []string {
	var headings []string
	scanner := bufio.NewScanner(bytes.NewReader(body))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ") {
			heading := strings.TrimSpace(strings.TrimPrefix(line, "## "))
			if heading != "" {
				headings = append(headings, heading)
			}
		}
	}
	return headings
}
