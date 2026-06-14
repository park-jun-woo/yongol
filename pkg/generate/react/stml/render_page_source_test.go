//ff:func feature=stml-gen type=test control=sequence topic=output
//ff:what renderPageSource — 최소 페이지에서 import·컴포넌트 함수 선언 골격 소스를 조립하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderPageSource(t *testing.T) {
	page := stmlparser.PageSpec{Name: "home"}
	out := renderPageSource(page, importSet{}, nil, nil, nil, "", "", GenerateOptions{})
	if !strings.Contains(out, "export default function") {
		t.Errorf("renderPageSource missing component declaration:\n%s", out)
	}
}
