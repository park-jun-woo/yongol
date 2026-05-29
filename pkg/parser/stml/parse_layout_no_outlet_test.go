//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLayoutNoOutlet — outlet 없는 레이아웃 파싱 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseLayoutNoOutlet(t *testing.T) {
	input := `<div>
  <nav>
    <a data-nav="/home">Home</a>
  </nav>
</div>`

	layout, diags := ParseLayoutReader("bare.html", "layouts/bare.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if layout.HasOutlet {
		t.Error("HasOutlet = true, want false")
	}
	if len(layout.NavItems) != 1 {
		t.Fatalf("NavItems = %d, want 1", len(layout.NavItems))
	}
	if layout.NavItems[0].Path != "/home" {
		t.Errorf("NavItems[0].Path = %q, want %q", layout.NavItems[0].Path, "/home")
	}
}
