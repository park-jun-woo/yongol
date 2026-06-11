//ff:func feature=stml-parse type=test control=sequence
//ff:what TestApplySitemapItemChild — li 자식 분기: 중첩 ul → Children, 첫 a href → 외부 링크, 비요소/href 없음/두 번째 a 무시 검증

package stml

import (
	"testing"

	"golang.org/x/net/html"
)

func TestApplySitemapItemChild(t *testing.T) {
	t.Run("nested ul appends children", func(t *testing.T) {
		li := firstElementNode(t, `<li>그룹<ul><li data-page="a">A</li><li data-page="b">B</li></ul></li>`, "li")
		var node SitemapNode
		var spec SitemapSpec
		for c := li.FirstChild; c != nil; c = c.NextSibling {
			applySitemapItemChild(c, &node, &spec)
		}
		if len(node.Children) != 2 || node.Children[0].Page != "a" || node.Children[1].Page != "b" {
			t.Errorf("Children = %+v, want pages a, b", node.Children)
		}
		if node.Href != "" {
			t.Errorf("Href = %q, want empty", node.Href)
		}
	})

	t.Run("first a href sets external link and fills empty label", func(t *testing.T) {
		li := firstElementNode(t, `<li><a href="https://docs.example.com">문서</a><a href="https://second.example.com">두 번째</a></li>`, "li")
		var node SitemapNode
		var spec SitemapSpec
		for c := li.FirstChild; c != nil; c = c.NextSibling {
			applySitemapItemChild(c, &node, &spec)
		}
		if node.Href != "https://docs.example.com" {
			t.Errorf("Href = %q, want first link kept", node.Href)
		}
		if node.Label != "문서" {
			t.Errorf("Label = %q, want 문서 from the link text", node.Label)
		}
	})

	t.Run("a does not overwrite an existing label", func(t *testing.T) {
		a := firstElementNode(t, `<li><a href="https://docs.example.com">링크 텍스트</a></li>`, "a")
		node := SitemapNode{Label: "직접 텍스트"}
		var spec SitemapSpec
		applySitemapItemChild(a, &node, &spec)
		if node.Href != "https://docs.example.com" || node.Label != "직접 텍스트" {
			t.Errorf("node = %+v, want href set and label preserved", node)
		}
	})

	t.Run("a without href is ignored", func(t *testing.T) {
		a := firstElementNode(t, `<li><a>이름 없는 링크</a></li>`, "a")
		var node SitemapNode
		var spec SitemapSpec
		applySitemapItemChild(a, &node, &spec)
		if node.Href != "" || node.Label != "" {
			t.Errorf("node = %+v, want untouched", node)
		}
	})

	t.Run("non-element and non-a children are ignored", func(t *testing.T) {
		var node SitemapNode
		var spec SitemapSpec
		applySitemapItemChild(&html.Node{Type: html.TextNode, Data: "텍스트"}, &node, &spec)
		span := firstElementNode(t, `<li><span>장식</span></li>`, "span")
		applySitemapItemChild(span, &node, &spec)
		if node.Href != "" || node.Label != "" || len(node.Children) != 0 {
			t.Errorf("node = %+v, want untouched", node)
		}
	})
}
