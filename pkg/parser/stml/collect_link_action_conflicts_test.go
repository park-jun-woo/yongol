//ff:func feature=stml-parse type=test control=sequence
//ff:what TestCollectLinkActionConflicts — data-link·data-action 동시 선언이 파스 ERROR로 차단됨을 검증

package stml

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCollectLinkActionConflicts(t *testing.T) {
	// Same element declaring both → parse-time ERROR.
	input := `<main>
  <button data-action="DeleteBuilding" data-link="building-detail">x</button>
</main>`
	_, diags := ParseReader("page.html", strings.NewReader(input))
	if len(diags) != 1 {
		t.Fatalf("got %d diags, want 1: %+v", len(diags), diags)
	}
	d := diags[0]
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseParse {
		t.Errorf("level/phase = %q/%q", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "data-link") || !strings.Contains(d.Message, "data-action") {
		t.Errorf("message = %q", d.Message)
	}
	if d.File != "page.html" {
		t.Errorf("file = %q", d.File)
	}

	// Separate elements → no conflict.
	ok := `<main>
  <a data-link="building-detail" data-link-params="route.BuildingID">상세</a>
  <button data-action="DeleteBuilding">삭제</button>
</main>`
	_, diags = ParseReader("page.html", strings.NewReader(ok))
	if len(diags) != 0 {
		t.Fatalf("expected no diags, got %+v", diags)
	}
}
