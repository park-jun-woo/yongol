//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what MenuRenderable — menuBlockReason 판정과의 일치 (렌더 true / 차단 false) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestMenuRenderable(t *testing.T) {
	node := stml.SitemapNode{Page: "building-list", Menu: true}

	t.Run("renderable node is true", func(t *testing.T) {
		if !MenuRenderable(node, 1, []string{"/buildings"}) {
			t.Error("expected true for a depth-1 page without required params")
		}
	})

	t.Run("required param is false", func(t *testing.T) {
		if MenuRenderable(node, 1, []string{"/buildings/:BuildingID"}) {
			t.Error("expected false for a required-param route")
		}
	})

	t.Run("depth 3 is false", func(t *testing.T) {
		if MenuRenderable(node, 3, []string{"/buildings"}) {
			t.Error("expected false beyond the 2-level render limit")
		}
	})

	t.Run("data-menu=false is false", func(t *testing.T) {
		hidden := stml.SitemapNode{Page: "building-list", Menu: false}
		if MenuRenderable(hidden, 1, []string{"/buildings"}) {
			t.Error("expected false for data-menu=\"false\"")
		}
	})
}
