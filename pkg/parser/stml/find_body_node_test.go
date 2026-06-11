//ff:func feature=stml-parse type=test control=sequence
//ff:what TestFindBodyNode — html.Parse 합성 트리에서 body 탐색 성공 / body 없는 수동 트리에서 nil 검증

package stml

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestFindBodyNode(t *testing.T) {
	t.Run("finds the synthesized body of a fragment", func(t *testing.T) {
		doc, err := html.Parse(strings.NewReader(`<nav data-sitemap></nav>`))
		if err != nil {
			t.Fatal(err)
		}
		body := findBodyNode(doc)
		if body == nil {
			t.Fatal("expected a body node, got nil")
		}
		if body.Type != html.ElementNode || body.Data != "body" {
			t.Errorf("node = type %v data %q, want body element", body.Type, body.Data)
		}
		if body.FirstChild == nil || body.FirstChild.Data != "nav" {
			t.Errorf("fragment top-level element should land under body, got %+v", body.FirstChild)
		}
	})

	t.Run("returns nil when no body exists", func(t *testing.T) {
		div := &html.Node{Type: html.ElementNode, Data: "div"}
		div.AppendChild(&html.Node{Type: html.ElementNode, Data: "span"})
		if body := findBodyNode(div); body != nil {
			t.Errorf("expected nil, got %+v", body)
		}
	})
}
