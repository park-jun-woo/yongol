//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what matchByGrep — 에러 메시지 키워드를 YAML 본문에서 grep하여 op 특정

package agent

import (
	"regexp"
	"strings"
)

var reQuotedKeyword = regexp.MustCompile("[\"'](\\w+)[\"']")

func matchByGrep(msg string, yamlContent string, offsets []pathOffset) ([]string, map[string]int) {
	keywords := reQuotedKeyword.FindAllStringSubmatch(msg, -1)
	if len(keywords) == 0 {
		return nil, nil
	}

	kwSet := make(map[string]bool, len(keywords))
	for _, m := range keywords {
		kwSet[m[1]] = true
	}

	lines := strings.Split(yamlContent, "\n")
	var hitLines []int
	for i, line := range lines {
		for kw := range kwSet {
			if strings.Contains(line, kw) {
				hitLines = append(hitLines, i+1)
				break
			}
		}
	}

	if len(hitLines) == 0 {
		return nil, nil
	}

	seen := make(map[string]bool)
	relativeLines := make(map[string]int)
	for _, lineNo := range hitLines {
		for _, off := range offsets {
			if lineNo >= off.StartLine && lineNo <= off.EndLine {
				if !seen[off.Op] {
					seen[off.Op] = true
					relativeLines[off.Op] = lineNo - off.StartLine
				}
				break
			}
		}
	}

	ops := make([]string, 0, len(seen))
	for op := range seen {
		ops = append(ops, op)
	}
	return ops, relativeLines
}
