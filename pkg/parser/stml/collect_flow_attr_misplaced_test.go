//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestCollectFlowAttrMisplaced — 허용 위치 밖 흐름 속성이 기록되고 정상 위치는 무시되는지 검증

package stml

import (
	"strings"
	"testing"
)

func TestCollectFlowAttrMisplaced(t *testing.T) {
	input := `<main>
  <div data-capture="a -> auth.token"><span>x</span></div>
  <div data-redirect="/x"><span>y</span></div>
  <p data-on-error>err</p>
  <section data-action="Login" data-capture="access_token -> auth.token" data-redirect="/">
    <p data-on-error></p>
    <button type="submit">go</button>
  </section>
</main>`

	page, _ := ParseReader("p.html", strings.NewReader(input))

	want := []FlowAttrMisplaced{
		{Attr: "data-capture", Tag: "div"},
		{Attr: "data-redirect", Tag: "div"},
		{Attr: "data-on-error", Tag: "p"},
	}
	if len(page.FlowAttrMisplaced) != len(want) {
		t.Fatalf("FlowAttrMisplaced = %+v, want %+v", page.FlowAttrMisplaced, want)
	}
	for i, w := range want {
		if page.FlowAttrMisplaced[i] != w {
			t.Errorf("[%d] = %+v, want %+v", i, page.FlowAttrMisplaced[i], w)
		}
	}

	// data-on-error on the data-action element itself is not "inside" the
	// block and is recorded as misplaced.
	input = `<main><form data-action="Login" data-on-error><button type="submit">go</button></form></main>`
	page, _ = ParseReader("p.html", strings.NewReader(input))
	if len(page.FlowAttrMisplaced) != 1 || page.FlowAttrMisplaced[0].Attr != "data-on-error" {
		t.Errorf("action root on-error: got %+v, want one data-on-error record", page.FlowAttrMisplaced)
	}
}
