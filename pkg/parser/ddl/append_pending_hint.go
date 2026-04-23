//ff:func feature=manifest type=parser control=sequence
//ff:what appendPendingHint — 주석 전용 라인에서 파싱된 힌트를 pending 에 누적

package ddl

// appendPendingHint parses a comment-only line and, when it yields a valid
// hint, appends it to the pending list for later attachment to the next DDL
// line.
func appendPendingHint(trim, path string, lineNum int, tableCtx string, pending []*HintComment) []*HintComment {
	hint := parseHintLine(trim, path, lineNum, tableCtx, "")
	if hint == nil {
		return pending
	}
	return append(pending, hint)
}
