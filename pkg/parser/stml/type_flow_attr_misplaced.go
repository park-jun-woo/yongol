//ff:type feature=stml-parse type=model
//ff:what FlowAttrMisplaced — 흐름 속성(data-capture/redirect/on-error)이 허용 위치 밖에 놓인 사실 기록
package stml

// FlowAttrMisplaced records a flow attribute found on an illegal position:
// data-capture / data-redirect on an element without data-action, or
// data-on-error outside any data-action block. The parser only records the
// fact; TM-25 turns it into an ERROR diagnostic at validate time.
type FlowAttrMisplaced struct {
	Attr string // offending attribute ("data-capture" | "data-redirect" | "data-on-error")
	Tag  string // HTML tag of the element carrying the attribute
}
