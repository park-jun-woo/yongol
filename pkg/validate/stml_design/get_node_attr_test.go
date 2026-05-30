//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestGetNodeAttr — getNodeAttr HTML 노드 속성 값 조회 분기 검증

package stml_design

import (
	"testing"

	"golang.org/x/net/html"
)

func TestGetNodeAttr(t *testing.T) {
	n := &html.Node{
		Attr: []html.Attribute{
			{Key: "class", Val: "bg-primary"},
			{Key: "id", Val: "main"},
		},
	}

	if got := getNodeAttr(n, "class"); got != "bg-primary" {
		t.Errorf("class = %q, want bg-primary", got)
	}
	if got := getNodeAttr(n, "id"); got != "main" {
		t.Errorf("id = %q, want main", got)
	}
	if got := getNodeAttr(n, "missing"); got != "" {
		t.Errorf("missing = %q, want empty", got)
	}
}
