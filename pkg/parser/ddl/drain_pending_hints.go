//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what drainPendingHints — pending 힌트를 실제 DDL 라인에 부착해 out 에 방출

package ddl

// drainPendingHints attaches every pending standalone hint to the current
// table context (and column when resolvable) and flushes them into out.
func drainPendingHints(pending []*HintComment, tableCtx, ddlTrim string, out []HintComment) ([]HintComment, []*HintComment) {
	column := extractColumnNameFromLine(ddlTrim)
	for _, h := range pending {
		if column != "" {
			h.ColumnCtx = column
		}
		h.TableCtx = tableCtx
		out = append(out, *h)
	}
	return out, nil
}
