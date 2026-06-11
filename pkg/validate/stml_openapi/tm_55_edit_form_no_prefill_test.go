//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-55 — GET-by-id + prefill 없는 PUT/PATCH 폼 WARNING, prefill 선언 시·GET-by-id 부재 시 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM55EditFormNoPrefill(t *testing.T) {
	opMap := buildOperationMethodMap(tm54Doc())

	t.Run("GET-by-id + no data-prefill on PUT fires WARNING", func(t *testing.T) {
		page := tm54Page(t, ``, `<input data-field="sheet_name" />`)
		diags := tm55EditFormNoPrefill(page, opMap)
		if countDiag(diags, "[TM-55]") != 1 || diags[0].Level != diagnostic.LevelWarning {
			t.Fatalf("expected 1 TM-55 WARNING, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "no data-prefill") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("declared data-prefill is silent", func(t *testing.T) {
		page := tm54Page(t, ` data-prefill="GetRule"`, `<input data-field="sheet_name" />`)
		if diags := tm55EditFormNoPrefill(page, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("no GET-by-id fetch on page is silent", func(t *testing.T) {
		// A bare PUT form without a route-param GET fetch: not the edit pattern.
		src := `<main><form data-action="UpdateRule"><input data-field="sheet_name" /><button type="submit">go</button></form></main>`
		page, diags := stml.ParseReader("p.html", strings.NewReader(src))
		if len(diags) > 0 {
			t.Fatal(diags)
		}
		if d := tm55EditFormNoPrefill(page, opMap); len(d) != 0 {
			t.Errorf("expected silence without a GET-by-id fetch, got %+v", d)
		}
	})
}
