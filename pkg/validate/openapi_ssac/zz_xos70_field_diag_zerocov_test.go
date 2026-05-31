//ff:func feature=validate type=test control=sequence
//ff:what TestXos70FieldDiag_ZeroCov — XOS-70 정수 응답 필드 format:int64 누락 진단 분기 직접 호출

package openapi_ssac

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXos70FieldDiag_ZeroCov(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetCount", FileName: "get_count.ssac"}

	// key absent in response constraints → (zero, false).
	if _, ok := xos70FieldDiag(fn, 1, "missing", "5", map[string]oapiparser.FieldConstraint{}); ok {
		t.Errorf("absent key should be false")
	}

	// non-integer field → false.
	rcStr := map[string]oapiparser.FieldConstraint{"name": {Type: "string"}}
	if _, ok := xos70FieldDiag(fn, 1, "name", `"x"`, rcStr); ok {
		t.Errorf("string field should be false")
	}

	// integer field, string literal (mismatch handled by XOS-67) → false.
	rcInt := map[string]oapiparser.FieldConstraint{"count": {Type: "integer"}}
	if _, ok := xos70FieldDiag(fn, 1, "count", `"oops"`, rcInt); ok {
		t.Errorf("string-to-integer should be skipped (XOS-67)")
	}

	// integer field, int literal, no format → diagnostic true.
	d, ok := xos70FieldDiag(fn, 3, "count", "10", rcInt)
	if !ok {
		t.Fatalf("int literal without format should yield diagnostic")
	}
	if d.Line != 3 || d.OperationID != "GetCount" {
		t.Errorf("diag = %+v", d)
	}

	// integer field, variable binding, no format → diagnostic true.
	if _, ok := xos70FieldDiag(fn, 4, "count", "total", rcInt); !ok {
		t.Errorf("variable binding should yield diagnostic")
	}

	// integer field already has format: int64 → false.
	rcFmt := map[string]oapiparser.FieldConstraint{"count": {Type: "integer", Format: "int64"}}
	if _, ok := xos70FieldDiag(fn, 5, "count", "10", rcFmt); ok {
		t.Errorf("format:int64 present should be false")
	}
}
