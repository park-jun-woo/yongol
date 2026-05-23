//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what splitSections — markdown 콘텐츠를 ## 헤딩별 섹션으로 분리

package agent

import "strings"

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
		}
	}
	if len(current) > 0 {
		sections = append(sections, strings.Join(current, "\n"))
	}
	return sections
}
