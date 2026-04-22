//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-4 테스트 — metadata.name negative

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestC04MetadataName_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Metadata: pmanifest.Metadata{Name: ""},
		},
	}
	got := c04MetadataName(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[C-4]") {
		t.Fatalf("message missing [C-4] prefix: %q", got[0].Message)
	}
}
