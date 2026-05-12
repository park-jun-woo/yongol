//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParsePageWithoutDataLayout — data-layout 속성 없는 페이지 검증

package stml

import (
	"strings"
	"testing"
)

func TestParsePageWithoutDataLayout(t *testing.T) {
	input := `<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">Login</button>
  </div>
</main>`

	page, diags := ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Layout != "" {
		t.Errorf("Layout = %q, want empty string", page.Layout)
	}
}
