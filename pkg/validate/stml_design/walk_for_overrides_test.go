//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestWalkForOverrides — walkForOverrides DOM 순회 @override class 추출 분기 검증
package stml_design

import (
	"testing"

	"golang.org/x/net/html"
)

func TestWalkForOverrides(t *testing.T) {
	root := &html.Node{Type: html.ElementNode, Data: "html"}
	// override comment with class value
	appendChild(root, &html.Node{Type: html.CommentNode, Data: `@override class="bg-foo p-8"`})
	// override comment without a class value -> cls == ""
	appendChild(root, &html.Node{Type: html.CommentNode, Data: `@override`})
	// non-override comment
	appendChild(root, &html.Node{Type: html.CommentNode, Data: `just text`})
	// nested element holding another override comment (exercises recursion)
	wrapper := &html.Node{Type: html.ElementNode, Data: "div"}
	appendChild(wrapper, &html.Node{Type: html.CommentNode, Data: `@override class="rounded-card"`})
	appendChild(root, wrapper)

	classes := make(map[string]bool)
	walkForOverrides(root, classes)

	if !classes["bg-foo p-8"] {
		t.Errorf("expected bg-foo p-8 captured, got %v", classes)
	}
	if !classes["rounded-card"] {
		t.Errorf("expected rounded-card captured via recursion, got %v", classes)
	}
	if len(classes) != 2 {
		t.Errorf("expected exactly 2 classes, got %v", classes)
	}
}
