//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-2 테스트 — APIVersion negative

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC02APIVersion_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{APIVersion: "yongol/v2"},
	}
	got := c02APIVersion(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[C-2]") {
		t.Fatalf("message missing [C-2] prefix: %q", got[0].Message)
	}
}
