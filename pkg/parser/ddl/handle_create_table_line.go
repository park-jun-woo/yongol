//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what handleCreateTableLine — CREATE TABLE 라인 감지 시 pending 힌트를 block-above 로 소비

package ddl

// handleCreateTableLine consumes any pending standalone hints as
// "above CREATE TABLE" and returns the updated table context, drained
// output slice and an empty pending list.
func handleCreateTableLine(trim string, pending []*HintComment, out []HintComment) (string, []HintComment, []*HintComment) {
	tableCtx := parseCreateTableName(trim)
	for _, h := range pending {
		h.TableCtx = tableCtx
		h.BlockAbove = true
		out = append(out, *h)
	}
	return tableCtx, out, nil
}
