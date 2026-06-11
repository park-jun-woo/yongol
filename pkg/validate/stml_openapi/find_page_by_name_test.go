//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestFindPageByName — 매칭 페이지의 슬라이스 내부 포인터 반환 / 미발견 nil 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFindPageByName(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}

	got := findPageByName(pages, "dashboard")
	if got == nil {
		t.Fatal("expected a page, got nil")
	}
	if got != &pages[1] {
		t.Errorf("expected a pointer into the slice (element 1), got %p", got)
	}
	if got.FileName != "dashboard.html" {
		t.Errorf("FileName = %q, want dashboard.html", got.FileName)
	}

	if got := findPageByName(pages, "ghost"); got != nil {
		t.Errorf("expected nil for an unknown name, got %+v", got)
	}
	if got := findPageByName(nil, "login"); got != nil {
		t.Errorf("expected nil for empty pages, got %+v", got)
	}
}
