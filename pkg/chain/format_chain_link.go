//ff:func feature=chain type=formatter control=selection
//ff:what chain link 한 줄을 SSOT/아티팩트 여부에 따라 포맷팅
package chain

import "fmt"

// formatChainLink formats a single chain link for display.
func formatChainLink(link Link, isArtifact bool) string {
	switch isArtifact {
	case true:
		loc := link.File
		if link.Summary != "" && link.Summary != "(file)" {
			loc = link.File + ":" + link.Summary
		}
		return fmt.Sprintf("  %-10s %-45s %s", link.Kind, loc, ownershipIcon(link.Ownership))
	default:
		loc := link.File
		if link.Line > 0 {
			loc = fmt.Sprintf("%s:%d", link.File, link.Line)
		}
		return fmt.Sprintf("  %-10s %-45s %s", link.Kind, loc, link.Summary)
	}
}
