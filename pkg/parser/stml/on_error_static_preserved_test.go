//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what 빈 data-on-error 마커 요소가 Children에 OnError StaticElement로 보존되는지 검증
package stml

import (
	"strings"
	"testing"
)

func TestParse_OnErrorStaticPreserved(t *testing.T) {
	page, err := ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
    <p class="error" data-on-error></p>
  </div>
</main>`))
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Actions) != 1 {
		t.Fatalf("expected 1 action, got %d", len(page.Actions))
	}
	a := page.Actions[0]
	if !a.OnErrorNode {
		t.Errorf("OnErrorNode = false, want true")
	}
	var found bool
	for _, ch := range a.Children {
		if ch.Kind != "static" || ch.Static == nil || !ch.Static.OnError {
			continue
		}
		found = true
		if ch.Static.Tag != "p" || ch.Static.ClassName != "error" {
			t.Errorf("on-error element = <%s class=%q>, want <p class=\"error\">", ch.Static.Tag, ch.Static.ClassName)
		}
	}
	if !found {
		t.Errorf("empty data-on-error element was not preserved in Children")
	}
}
