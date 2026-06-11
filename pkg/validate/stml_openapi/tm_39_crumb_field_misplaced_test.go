//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm39CrumbFieldMisplaced — 그룹 항목의 data-crumb-field ERROR 발화 / 속성 없으면 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM39CrumbFieldMisplaced(t *testing.T) {
	t.Run("group entry with data-crumb-field fires", func(t *testing.T) {
		e := sitemapEntry{Node: stml.SitemapNode{Label: "건물", CrumbField: "name"}, Path: "nav[0] > 건물"}
		diags := tm39CrumbFieldMisplaced(e, "sitemap.html")
		if len(diags) != 1 || diags[0].Level != diagnostic.LevelError {
			t.Fatalf("expected 1 ERROR, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "[TM-39]") || !strings.Contains(diags[0].Message, "nav[0] > 건물") {
			t.Errorf("Message should carry the rule ID and position, got %q", diags[0].Message)
		}
	})

	t.Run("group entry without the attribute is silent", func(t *testing.T) {
		e := sitemapEntry{Node: stml.SitemapNode{Label: "건물"}, Path: "nav[0] > 건물"}
		if diags := tm39CrumbFieldMisplaced(e, "sitemap.html"); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
