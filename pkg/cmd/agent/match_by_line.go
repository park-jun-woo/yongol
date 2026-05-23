//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what matchByLine — YAML line 에러에서 라인 번호 추출 후 op 매핑

package agent

import (
	"regexp"
	"strconv"
)

var reYAMLLine = regexp.MustCompile("(?:yaml: )?line (\\d+)")

func matchByLine(msg string, offsets []pathOffset) []string {
	matches := reYAMLLine.FindAllStringSubmatch(msg, -1)
	if len(matches) == 0 {
		return nil
	}
	var ops []string
	for _, m := range matches {
		lineNo, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}
		for _, off := range offsets {
			if lineNo >= off.StartLine && lineNo <= off.EndLine {
				ops = append(ops, off.Op)
				break
			}
		}
	}
	return ops
}
