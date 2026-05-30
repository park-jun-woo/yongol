//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestIsPrecededByOverride — isPrecededByOverride @override 주석 선행 여부 분기 검증

package stml_design

import (
	"testing"

	"golang.org/x/net/html"
)

// link wires prev as the immediate previous sibling of n.
func link(prev, n *html.Node) {
	n.PrevSibling = prev
}

func TestIsPrecededByOverride(t *testing.T) {
	t.Run("no previous sibling", func(t *testing.T) {
		n := &html.Node{Type: html.ElementNode, Data: "div"}
		if isPrecededByOverride(n) {
			t.Error("expected false with no previous sibling")
		}
	})

	t.Run("immediate override comment", func(t *testing.T) {
		comment := &html.Node{Type: html.CommentNode, Data: "@override class=bg-foo"}
		n := &html.Node{Type: html.ElementNode, Data: "div"}
		link(comment, n)
		if !isPrecededByOverride(n) {
			t.Error("expected true with override comment")
		}
	})

	t.Run("whitespace text then override comment", func(t *testing.T) {
		comment := &html.Node{Type: html.CommentNode, Data: "@override"}
		ws := &html.Node{Type: html.TextNode, Data: "\n  "}
		n := &html.Node{Type: html.ElementNode, Data: "div"}
		ws.PrevSibling = comment
		link(ws, n)
		if !isPrecededByOverride(n) {
			t.Error("expected true skipping whitespace")
		}
	})

	t.Run("non-whitespace text breaks", func(t *testing.T) {
		text := &html.Node{Type: html.TextNode, Data: "hello"}
		n := &html.Node{Type: html.ElementNode, Data: "div"}
		link(text, n)
		if isPrecededByOverride(n) {
			t.Error("expected false when preceded by real text")
		}
	})

	t.Run("non-override comment breaks", func(t *testing.T) {
		comment := &html.Node{Type: html.CommentNode, Data: "just a comment"}
		n := &html.Node{Type: html.ElementNode, Data: "div"}
		link(comment, n)
		if isPrecededByOverride(n) {
			t.Error("expected false for non-override comment")
		}
	})
}
