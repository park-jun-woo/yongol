//ff:func feature=rule type=parser control=selection topic=catalog
//ff:what rulebookParseState.feedTableRow — `| ... |` 테이블 라인을 header/separator/data 로 분기

package catalog

func (st *rulebookParseState) feedTableRow(trimmed string) {
	switch {
	case !st.sawHeader:
		if isRuleTableHeader(trimmed) {
			st.sawHeader = true
		}
	case !st.inTable:
		if isTableSeparator(trimmed) {
			st.inTable = true
		} else {
			// Unexpected row between header and separator — bail on this table.
			st.sawHeader = false
		}
	default:
		st.appendDataRow(trimmed)
	}
}
