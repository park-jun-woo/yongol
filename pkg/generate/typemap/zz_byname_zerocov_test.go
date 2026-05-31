//ff:func feature=gen-typemap type=test control=sequence
//ff:what typemap 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package typemap

import "testing"

func TestClassifyNativeFamily_ZeroCov(t *testing.T) {
	if fam, ok := classifyNativeFamily("BIGINT"); !ok || fam != FamilyInteger {
		t.Errorf("BIGINT -> %v ok=%v", fam, ok)
	}
	if fam, ok := classifyNativeFamily("BOOLEAN"); !ok || fam != FamilyBoolean {
		t.Errorf("BOOLEAN -> %v ok=%v", fam, ok)
	}
	if _, ok := classifyNativeFamily("NOPE"); ok {
		t.Error("unknown should be false")
	}
}

func TestClassifyPgtypeFamily_ZeroCov(t *testing.T) {
	if fam, ok := classifyPgtypeFamily("UUID"); !ok || fam != FamilyUUID {
		t.Errorf("UUID -> %v ok=%v", fam, ok)
	}
	if fam, ok := classifyPgtypeFamily("JSONB"); !ok || fam != FamilyJSONB {
		t.Errorf("JSONB -> %v ok=%v", fam, ok)
	}
	if _, ok := classifyPgtypeFamily("NOPE"); ok {
		t.Error("unknown should be false")
	}
}

func TestParseRawType_ZeroCov(t *testing.T) {
	info := ParseRawType("VARCHAR(255)[]")
	if !info.IsArray || info.Param != "255" || info.Head != "VARCHAR" {
		t.Errorf("array/param/head wrong: %#v", info)
	}
	multi := ParseRawType("DOUBLE PRECISION")
	if !multi.MultiToken {
		t.Errorf("expected multi-token: %#v", multi)
	}
}
