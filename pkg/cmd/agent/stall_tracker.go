//ff:type feature=agent type=helper
//ff:what stallTracker — 파일별 연속 동일 진단 추적

package agent

// stallTracker tracks consecutive identical diagnostics for a file.
type stallTracker struct {
	lastMessages string
	count        int
}
