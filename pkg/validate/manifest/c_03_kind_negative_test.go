//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-3 테스트 — Kind negative

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC03Kind_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Kind: ""},
	}
	got := c03Kind(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[C-3]") {
		t.Fatalf("message missing [C-3] prefix: %q", got[0].Message)
	}
}
