//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-54 — prefill 소스 부재(ERROR)·필드 커버리지(WARNING)·정상 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestTM54PrefillSource(t *testing.T) {
	opMap := buildOperationMethodMap(tm54Doc())

	t.Run("unknown prefill source fires ERROR", func(t *testing.T) {
		page := tm54Page(t, ` data-prefill="NoSuchFetch"`, `<input data-field="sheet_name" />`)
		diags := tm54PrefillSource(page, opMap)
		if countDiag(diags, "[TM-54]") != 1 || diags[0].Level != diagnostic.LevelError {
			t.Fatalf("expected 1 TM-54 ERROR, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "out of scope") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("field absent from prefill response fires WARNING", func(t *testing.T) {
		page := tm54Page(t, ` data-prefill="GetRule"`, `<input data-field="sheet_name" /><input data-field="note" />`)
		diags := tm54PrefillSource(page, opMap)
		// "note" is not in GetRule's response.
		if countDiag(diags, "[TM-54]") != 1 || diags[0].Level != diagnostic.LevelWarning {
			t.Fatalf("expected 1 TM-54 WARNING, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "note") || !strings.Contains(diags[0].Message, "opens blank") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("all fields covered is silent", func(t *testing.T) {
		page := tm54Page(t, ` data-prefill="GetRule"`, `<input data-field="sheet_name" /><input data-field="start_row" />`)
		if diags := tm54PrefillSource(page, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("no data-prefill is silent", func(t *testing.T) {
		page := tm54Page(t, ``, `<input data-field="sheet_name" />`)
		if diags := tm54PrefillSource(page, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
