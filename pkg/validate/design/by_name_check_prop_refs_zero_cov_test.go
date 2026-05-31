//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — design 토큰 참조/미지 prop 검사 헬퍼 직접 호출
package design

import (
	"testing"
)

func TestByNameCheckPropRefs_ZeroCov(t *testing.T) {
	fs := designFS()
	props := map[string]string{
		"color":  "{colors.primary}", // resolves
		"border": "{colors.nope}",    // unresolved → diag
	}
	diags := checkPropRefs(fs, "Button", props)
	if len(diags) != 1 {
		t.Errorf("expected 1 unresolved-ref diag, got %d", len(diags))
	}

	// single prop helper directly: resolved → no diag.
	if d := checkSinglePropRefs(fs, "Button", "color", "{colors.primary}"); len(d) != 0 {
		t.Errorf("resolved single prop should yield no diag, got %v", d)
	}
	// no refs at all.
	if d := checkSinglePropRefs(fs, "Button", "label", "Save"); len(d) != 0 {
		t.Errorf("no-ref value should yield no diag")
	}
}
