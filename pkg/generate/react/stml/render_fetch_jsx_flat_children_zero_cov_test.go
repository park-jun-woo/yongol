//ff:func feature=stml-gen type=test control=sequence
//ff:what zz_zerocov_test — renderFetchJSXFlatChildren / renderPageJSXFallback 0% 커버리지 단위 테스트
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderFetchJSXFlatChildren_ZeroCov(t *testing.T) {
	f := stmlparser.FetchBlock{
		OperationID: "ListItems",
		Binds:       []stmlparser.FieldBind{{Name: "name", Tag: "span"}},
		Eaches:      []stmlparser.EachBlock{{Tag: "ul", ItemTag: "li"}},
		States:      []stmlparser.StateBind{{Tag: "p"}},
		Components:  []stmlparser.ComponentRef{{Name: "DatePicker"}},
	}
	lines := renderFetchJSXFlatChildren(f, "data", 0, map[string]bool{})
	// 1 bind + 1 each + 1 state + 1 component = 4 lines.
	if len(lines) != 4 {
		t.Fatalf("expected 4 lines, got %d: %v", len(lines), lines)
	}

	// Empty fetch → no lines.
	if got := renderFetchJSXFlatChildren(stmlparser.FetchBlock{}, "data", 0, nil); len(got) != 0 {
		t.Errorf("empty fetch should yield 0 lines, got %v", got)
	}
}
