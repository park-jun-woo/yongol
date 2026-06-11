//ff:func feature=stml-gen type=test control=sequence
//ff:what zodPrefillValues — 전 필드 방출·응답 교집합 data 참조·나머지 빈 리터럴·정렬 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestZodPrefillValues(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"sheet_name": {Type: "string"},
		"start_row":  {Type: "integer"},
		"note":       {Type: "string"},
	}
	resp := map[string]oapiparser.FieldTypeInfo{"sheet_name": {Type: "string"}, "start_row": {Type: "integer"}}
	out := zodPrefillValues(fields, "getRuleData", resp)

	if !strings.Contains(out, "sheet_name: getRuleData.sheet_name ?? '',") {
		t.Errorf("present string field: %q", out)
	}
	if !strings.Contains(out, "start_row: getRuleData.start_row ?? 0,") {
		t.Errorf("present integer field: %q", out)
	}
	if !strings.Contains(out, "note: '',") || strings.Contains(out, "getRuleData.note") {
		t.Errorf("absent field must use empty literal, not data: %q", out)
	}
	// sorted: note, sheet_name, start_row
	if i, j := strings.Index(out, "note:"), strings.Index(out, "sheet_name:"); i > j {
		t.Errorf("fields not sorted: %q", out)
	}
}
