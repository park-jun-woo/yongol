//ff:func feature=gen-typemap type=test-helper control=sequence
//ff:what assertRawType — ParseRawType 결과 필드별 일치 검증 헬퍼

package typemap

import "testing"

func assertRawType(t *testing.T, raw, wantHead, wantParam string, wantArray, wantMulti bool) {
	t.Helper()
	got := ParseRawType(raw)
	if got.Head != wantHead {
		t.Errorf("Head = %q, want %q", got.Head, wantHead)
	}
	if got.Param != wantParam {
		t.Errorf("Param = %q, want %q", got.Param, wantParam)
	}
	if got.IsArray != wantArray {
		t.Errorf("IsArray = %v, want %v", got.IsArray, wantArray)
	}
	if got.MultiToken != wantMulti {
		t.Errorf("MultiToken = %v, want %v", got.MultiToken, wantMulti)
	}
}
