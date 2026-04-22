//ff:func feature=gen-gogin type=util control=sequence
//ff:what scanLine — 한 줄을 (depth-0 제어구조 종류, depth 변화량)으로 환산

package ffannot

import "strings"

// scanLine returns (controlKind, braceDelta) for a single body line.
// When depth == 0 and the trimmed line opens a for/switch, controlKind is set
// to ControlIteration or ControlSelection; otherwise "". braceDelta is the
// net change in { } depth contributed by this line (comments stripped).
// Empty/whitespace lines produce ("", 0).
func scanLine(raw string, depth int) (string, int) {
	line := strings.TrimSpace(raw)
	if line == "" {
		return "", 0
	}
	var kind string
	if depth == 0 {
		kind = classifyDepth0(line)
	}
	return kind, countBraceDelta(line)
}
