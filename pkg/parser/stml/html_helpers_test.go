//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText

package stml

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// firstElementNode parses an HTML fragment and returns the first element
// node whose tag matches `tag`.
func firstElementNode(t *testing.T, fragment, tag string) *html.Node {
	t.Helper()
	doc, err := html.Parse(strings.NewReader(fragment))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var found *html.Node
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if found != nil {
			return
		}
		if n.Type == html.ElementNode && n.Data == tag {
			found = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	if found == nil {
		t.Fatalf("tag %q not found in %q", tag, fragment)
	}
	return found
}

func TestIsImplicitTag(t *testing.T) {
	for _, tag := range []string{"html", "head", "body"} {
		if !isImplicitTag(tag) {
			t.Errorf("%q should be implicit", tag)
		}
	}
	if isImplicitTag("div") {
		t.Errorf("div should not be implicit")
	}
}

func TestGetAttrAndHasAttr(t *testing.T) {
	n := firstElementNode(t, `<div id="x" data-fetch="/api"></div>`, "div")
	if got := getAttr(n, "id"); got != "x" {
		t.Errorf("getAttr id = %q, want x", got)
	}
	if got := getAttr(n, "missing"); got != "" {
		t.Errorf("getAttr missing = %q, want empty", got)
	}
	if !hasAttr(n, "data-fetch") {
		t.Errorf("hasAttr data-fetch should be true")
	}
	if hasAttr(n, "nope") {
		t.Errorf("hasAttr nope should be false")
	}
}

func TestHasDataAttr(t *testing.T) {
	withData := firstElementNode(t, `<div data-x="1"></div>`, "div")
	if !hasDataAttr(withData) {
		t.Errorf("expected data attr")
	}
	without := firstElementNode(t, `<div class="c"></div>`, "div")
	if hasDataAttr(without) {
		t.Errorf("expected no data attr")
	}
}

func TestIsDataElement(t *testing.T) {
	fetch := firstElementNode(t, `<div data-fetch="/a"></div>`, "div")
	if !isDataElement(fetch) {
		t.Errorf("data-fetch should be data element")
	}
	action := firstElementNode(t, `<form data-action="/a"></form>`, "form")
	if !isDataElement(action) {
		t.Errorf("data-action should be data element")
	}
	plain := firstElementNode(t, `<div class="c"></div>`, "div")
	if isDataElement(plain) {
		t.Errorf("plain div should not be data element")
	}
}

func TestHasFieldAttr(t *testing.T) {
	field := firstElementNode(t, `<span data-field="name"></span>`, "span")
	if !hasFieldAttr(field) {
		t.Errorf("data-field should match")
	}
	none := firstElementNode(t, `<span data-component="c"></span>`, "span")
	if hasFieldAttr(none) {
		t.Errorf("data-component alone should not match")
	}
}

func TestSplitTrim(t *testing.T) {
	got := splitTrim("a, b ,, c ")
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if g := splitTrim("  ,  "); g != nil {
		t.Errorf("all-empty = %v, want nil", g)
	}
}

func TestExtractAllTextAndCollectText(t *testing.T) {
	n := firstElementNode(t, `<div>Hello <b>World</b>!</div>`, "div")
	if got := extractAllText(n); got != "Hello World!" {
		t.Errorf("extractAllText = %q, want %q", got, "Hello World!")
	}
	var sb strings.Builder
	collectText(n, &sb)
	if !strings.Contains(sb.String(), "World") {
		t.Errorf("collectText missing World: %q", sb.String())
	}
}

func TestDirectTextAndExtractNonEmptyText(t *testing.T) {
	n := firstElementNode(t, `<div>   <span>x</span>label</div>`, "div")
	// directText returns first non-empty direct text child; whitespace-only
	// text node is skipped, then the element, then "label".
	if got := directText(n); got != "label" {
		t.Errorf("directText = %q, want label", got)
	}
	// extractNonEmptyText on a non-text node returns ""
	span := firstElementNode(t, `<span>y</span>`, "span")
	if got := extractNonEmptyText(span); got != "" {
		t.Errorf("extractNonEmptyText(element) = %q, want empty", got)
	}
	if got := extractNonEmptyText(span.FirstChild); got != "y" {
		t.Errorf("extractNonEmptyText(text) = %q, want y", got)
	}
}

func TestHasContent(t *testing.T) {
	withText := firstElementNode(t, `<div>hello</div>`, "div")
	if !hasContent(withText) {
		t.Errorf("text content should count")
	}
	withChild := firstElementNode(t, `<div><span></span></div>`, "div")
	if !hasContent(withChild) {
		t.Errorf("element child should count")
	}
	empty := firstElementNode(t, `<div>   </div>`, "div")
	if hasContent(empty) {
		t.Errorf("whitespace-only should not count")
	}
}
