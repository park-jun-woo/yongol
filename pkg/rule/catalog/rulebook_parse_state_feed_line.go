//ff:func feature=rule type=parser control=selection topic=catalog
//ff:what rulebookParseState.feedLine — 한 줄을 읽어 섹션/테이블/데이터 라인으로 분기 처리

package catalog

import "strings"

func (st *rulebookParseState) feedLine(line string) {
	trimmed := strings.TrimSpace(line)

	switch {
	case strings.HasPrefix(trimmed, "## "):
		st.handleSectionHeading(trimmed)
	case strings.HasPrefix(trimmed, "### "), strings.HasPrefix(trimmed, "#### "):
		st.resetTable()
	case st.sectionSkip:
		// skip everything inside ## Deprecated
	case trimmed == "":
		st.resetTable()
	case !strings.HasPrefix(trimmed, "|"):
		st.resetTable()
	default:
		st.feedTableRow(trimmed)
	}
}
