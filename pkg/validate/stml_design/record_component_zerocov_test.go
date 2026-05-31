//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestRecordComponent_ZeroCov — recordComponent 양/음 분기 커버

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRecordComponent_ZeroCov(t *testing.T) {
	out := &pageTokenRefs{}
	// empty name → not recorded
	recordComponent(stml.ComponentRef{Name: ""}, "page.stml", out)
	if len(out.Components) != 0 {
		t.Fatalf("empty name should not record, got %d", len(out.Components))
	}
	// named → recorded
	recordComponent(stml.ComponentRef{Name: "DatePicker"}, "page.stml", out)
	if len(out.Components) != 1 || out.Components[0].Name != "DatePicker" || out.Components[0].File != "page.stml" {
		t.Fatalf("recordComponent = %+v", out.Components)
	}
}
