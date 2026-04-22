//ff:func feature=gen-gogin type=util control=selection
//ff:what classifyDepth0 — depth-0 한 줄을 iteration/selection/""(none)으로 분류

package ffannot

import "strings"

// classifyDepth0 reports whether a depth-0 line is a loop or switch header.
// Returns ControlIteration, ControlSelection, or "" when neither.
func classifyDepth0(line string) string {
	switch {
	case strings.HasPrefix(line, "for ") ||
		strings.HasPrefix(line, "for{") ||
		strings.HasPrefix(line, "for(") ||
		line == "for":
		return ControlIteration
	case strings.HasPrefix(line, "switch ") ||
		strings.HasPrefix(line, "switch{") ||
		line == "switch":
		return ControlSelection
	}
	return ""
}
