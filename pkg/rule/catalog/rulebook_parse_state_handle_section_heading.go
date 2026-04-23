//ff:func feature=rule type=parser control=sequence topic=catalog
//ff:what rulebookParseState.handleSectionHeading — `## <title>` 줄을 섹션 상태에 반영

package catalog

import "strings"

func (st *rulebookParseState) handleSectionHeading(trimmed string) {
	st.sectionTitle = strings.TrimSpace(strings.TrimPrefix(trimmed, "##"))
	st.sectionSkip = strings.EqualFold(st.sectionTitle, "Deprecated")
	st.inTable = false
	st.sawHeader = false
}
