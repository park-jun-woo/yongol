//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParsePageWithoutDataRoute — data-route 속성 없는 페이지 검증

package stml

import (
	"strings"
	"testing"
)

func TestParsePageWithoutDataRoute(t *testing.T) {
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
	if page.Route != "" {
		t.Errorf("Route = %q, want empty string", page.Route)
	}
}
