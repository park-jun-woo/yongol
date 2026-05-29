//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-6 테스트 — backend zero 값일 때도 auth nil 이므로 ERROR 1건

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestC06BackendAuthRequired_BackendAbsent_Error — backend 자체가 zero 값일
// 때도 auth 가 nil 이므로 [C-6] ERROR. (병행하는 [C-5] 도 같이 fire 가능 —
// 본 룰은 독립적으로 검사.)
func TestC06BackendAuthRequired_BackendAbsent_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
	}
	got := c06BackendAuthRequired(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[C-6]") {
		t.Fatalf("message missing [C-6] prefix: %q", got[0].Message)
	}
}
