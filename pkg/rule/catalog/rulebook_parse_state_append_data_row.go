//ff:func feature=rule type=parser control=sequence topic=catalog
//ff:what rulebookParseState.appendDataRow — 데이터 라인을 파싱해 RuleMeta 추가 (잘못된 행은 무시)

package catalog

import "strings"

func (st *rulebookParseState) appendDataRow(trimmed string) {
	cells := splitRow(trimmed)
	if len(cells) < 4 {
		return
	}
	id := strings.TrimSpace(cells[0])
	level := strings.ToUpper(strings.TrimSpace(cells[1]))
	desc := strings.TrimSpace(cells[2])
	source := strings.Trim(strings.TrimSpace(cells[3]), "`")

	if id == "" || (level != "ERROR" && level != "WARNING") {
		return
	}
	st.rules = append(st.rules, RuleMeta{
		ID:            id,
		Level:         level,
		Description:   desc,
		Source:        source,
		SectionTitle:  st.sectionTitle,
		SectionAnchor: sectionAnchor(st.sectionTitle),
	})
}
