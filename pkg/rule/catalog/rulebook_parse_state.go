//ff:type feature=rule type=model topic=catalog
//ff:what rulebookParseState — Parse 루프 1회에서 유지해야 하는 섹션/테이블 상태 집합

package catalog

type rulebookParseState struct {
	rules        []RuleMeta
	sectionTitle string
	sectionSkip  bool // true while inside ## Deprecated
	inTable      bool // past header + separator rows
	sawHeader    bool // saw "| Rule ID | Level | ..." row
}
