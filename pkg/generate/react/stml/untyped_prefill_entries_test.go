//ff:func feature=stml-gen type=test control=sequence
//ff:what untypedPrefillEntries — 응답 교집합 필드만 정렬·중복제거 data 참조, 교집합 없으면 빈 문자열 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUntypedPrefillEntries(t *testing.T) {
	fields := []stmlparser.FieldBind{{Name: "b"}, {Name: "a"}, {Name: "a"}, {Name: ""}, {Name: "z"}}
	resp := map[string]oapiparser.FieldTypeInfo{"a": {Type: "string"}, "b": {Type: "string"}}
	out := untypedPrefillEntries(fields, "d", resp)
	if !strings.Contains(out, "a: d.a,") || !strings.Contains(out, "b: d.b,") {
		t.Errorf("present fields: %q", out)
	}
	if strings.Contains(out, "d.z") {
		t.Errorf("absent field z must be dropped: %q", out)
	}
	if strings.Count(out, "a: d.a,") != 1 {
		t.Errorf("duplicate field a not deduped: %q", out)
	}
	if i, j := strings.Index(out, "a:"), strings.Index(out, "b:"); i > j {
		t.Errorf("not sorted: %q", out)
	}

	// No overlap → empty string.
	if got := untypedPrefillEntries(fields, "d", map[string]oapiparser.FieldTypeInfo{"x": {}}); got != "" {
		t.Errorf("no overlap should yield empty, got %q", got)
	}
}
