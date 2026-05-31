//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestWalkInlineStyles — walkInlineStyles DOM 순회 inline style 색상 검사 분기 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

func TestWalkInlineStyles(t *testing.T) {
	hexToToken := map[string]string{"#6366f1": "primary"}

	root := &html.Node{Type: html.ElementNode, Data: "html"}

	// element with hardcoded token color -> fires
	styled := &html.Node{
		Type: html.ElementNode, Data: "div",
		Attr: []html.Attribute{{Key: "style", Val: "color: #6366F1"}},
	}
	appendChild(root, styled)

	// element with no style -> skipped (covers style == "" guard)
	plain := &html.Node{Type: html.ElementNode, Data: "span"}
	appendChild(root, plain)

	// element preceded by @override comment -> skipped
	overrideComment := &html.Node{Type: html.CommentNode, Data: `@override class="x"`}
	overridden := &html.Node{
		Type: html.ElementNode, Data: "p",
		Attr: []html.Attribute{{Key: "style", Val: "color: #6366F1"}},
	}
	overridden.PrevSibling = overrideComment
	appendChild(root, overrideComment)
	appendChild(root, overridden)

	var diags []diagnostic.Diagnostic
	walkInlineStyles(root, "page.html", hexToToken, overrideSet{}, &diags)

	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic (only the non-overridden styled element), got %d: %+v", len(diags), diags)
	}
}
