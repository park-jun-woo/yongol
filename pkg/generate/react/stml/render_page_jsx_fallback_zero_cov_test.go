//ff:func feature=stml-gen type=test control=sequence
//ff:what zz_zerocov_test — renderFetchJSXFlatChildren / renderPageJSXFallback 0% 커버리지 단위 테스트
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderPageJSXFallback_ZeroCov(t *testing.T) {
	page := stmlparser.PageSpec{
		Fetches: []stmlparser.FetchBlock{{OperationID: "ListItems", Tag: "section"}},
		Actions: []stmlparser.ActionBlock{{OperationID: "CreateItem", Tag: "form"}},
	}
	var sb strings.Builder
	renderPageJSXFallback(page, &sb, map[string]bool{})
	out := sb.String()
	if out == "" {
		t.Fatal("expected non-empty output")
	}

	// Empty page → empty output.
	var sb2 strings.Builder
	renderPageJSXFallback(stmlparser.PageSpec{}, &sb2, nil)
	if sb2.String() != "" {
		t.Errorf("empty page should yield empty output, got %q", sb2.String())
	}
}
