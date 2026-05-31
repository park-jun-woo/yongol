//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameHandlers_ZeroCov(t *testing.T) {
	compEl := firstElementNode(t, `<span data-component="Card" data-bind="b" data-field="f"></span>`, "span")
	var ab ActionBlock
	handleActionComponent(compEl, &ab)
	handleStaticActionComponent(compEl, &ab)

	fieldEl := firstElementNode(t, `<input data-field="Name" type="text"/>`, "input")
	handleActionField(fieldEl, &ab)
	handleStaticActionField(fieldEl, &ab)

	var fb FetchBlock
	handleFetchComponent(compEl, &fb)

	// walk-static action helpers
	se := StaticElement{Tag: "div"}
	var ab2 ActionBlock
	handleWalkStaticActionComponent(compEl, &ab2, &se)
	handleWalkStaticActionField(fieldEl, &ab2, &se)
}
