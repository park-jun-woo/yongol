//ff:func feature=stml-gen type=test control=sequence
//ff:what TestMarkLinkImports — 단일 LinkRef의 Link/useParams 플래그 설정 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestMarkLinkImports(t *testing.T) {
	var is importSet
	markLinkImports(stmlparser.LinkRef{Params: []stmlparser.LinkParamBind{{Source: "item.id"}}}, &is)
	if !is.useLink || is.useParams {
		t.Errorf("item source: useLink=%v useParams=%v", is.useLink, is.useParams)
	}
	var is2 importSet
	markLinkImports(stmlparser.LinkRef{Params: []stmlparser.LinkParamBind{{Source: "route.ID"}}}, &is2)
	if !is2.useLink || !is2.useParams {
		t.Errorf("route source: useLink=%v useParams=%v", is2.useLink, is2.useParams)
	}
}
