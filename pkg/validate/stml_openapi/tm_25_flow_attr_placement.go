//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-25 — data-on-error가 data-action 블록 밖, 또는 data-capture/redirect가 data-action 없는 요소에 위치 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm25FlowAttrPlacement turns the parser's flow-attribute misplacement
// records into ERROR diagnostics: data-capture and data-redirect belong on
// the data-action element itself, and data-on-error belongs on an element
// inside a data-action block. A misplaced flow attribute would silently do
// nothing at runtime, so it is rejected at validate time.
func tm25FlowAttrPlacement(page stml.PageSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, m := range page.FlowAttrMisplaced {
		var msg, advice string
		if m.Attr == "data-on-error" {
			msg = fmt.Sprintf("[TM-25] data-on-error on <%s> is outside any data-action block", m.Tag)
			advice = "Move the data-on-error element inside a data-action block; it shows the server error message when that action fails"
		} else {
			msg = fmt.Sprintf("[TM-25] %s on <%s> requires data-action on the same element", m.Attr, m.Tag)
			advice = fmt.Sprintf("Move %s onto the element that declares data-action, or add data-action to this element", m.Attr)
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    page.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: msg,
			Advice:  advice,
		})
	}
	return diags
}
