//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-56 — PATCH 전필드 required 폼 WARNING, 일부 optional·PUT·미소비 시 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestTM56PatchAllRequired(t *testing.T) {
	t.Run("PATCH with all fields required fires WARNING", func(t *testing.T) {
		opMap := buildOperationMethodMap(tm56Doc([]string{"sheet_name", "start_row"}))
		diags := tm56PatchAllRequired(tm56Page(t), opMap)
		if countDiag(diags, "[TM-56]") != 1 || diags[0].Level != diagnostic.LevelWarning {
			t.Fatalf("expected 1 TM-56 WARNING, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "all-required") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("PATCH with an optional field is silent", func(t *testing.T) {
		opMap := buildOperationMethodMap(tm56Doc([]string{"sheet_name"}))
		if diags := tm56PatchAllRequired(tm56Page(t), opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("PUT all-required is silent (only PATCH is scoped)", func(t *testing.T) {
		opMap := buildOperationMethodMap(tm54Doc())
		page := tm54Page(t, ``, `<input data-field="sheet_name" />`)
		if diags := tm56PatchAllRequired(page, opMap); len(diags) != 0 {
			t.Errorf("expected silence for PUT, got %+v", diags)
		}
	})
}
