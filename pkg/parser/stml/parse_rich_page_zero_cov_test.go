//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseRichPage_ZeroCov — 정적 래퍼/필드/컴포넌트/each/action 분기 전부 도달
package stml

import (
	"strings"
	"testing"
)

func TestParseRichPage_ZeroCov(t *testing.T) {
	page, diags := ParseReader("things.html", strings.NewReader(richPageHTML))
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if page.Route != "/things/:id" {
		t.Errorf("Route = %q, want /things/:id", page.Route)
	}
	if len(page.Fetches) == 0 {
		t.Fatalf("expected at least one fetch block")
	}
	fb := page.Fetches[0]
	if fb.OperationID != "GetThing" {
		t.Errorf("fetch OperationID = %q, want GetThing", fb.OperationID)
	}
	// Component collected via handleFetchComponent.
	foundComp := false
	for _, c := range fb.Components {
		if c.Name == "Avatar" {
			foundComp = true
		}
	}
	if !foundComp {
		t.Errorf("expected Avatar component collected, got %+v", fb.Components)
	}
	if len(fb.Eaches) == 0 {
		t.Errorf("expected each block collected")
	}
}
