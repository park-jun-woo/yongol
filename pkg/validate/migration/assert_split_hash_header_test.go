//ff:func feature=validate type=test-helper control=sequence
//ff:what assertSplitHashHeader — splitHashHeader 결과 ok/head/body 검증 헬퍼

package migration

import "testing"

func assertSplitHashHeader(t *testing.T, text string, wantOK bool, wantHead, wantBody string) {
	t.Helper()
	head, body, ok := splitHashHeader(text)
	if ok != wantOK {
		t.Fatalf("ok = %v, want %v", ok, wantOK)
	}
	if !wantOK {
		return
	}
	if head != wantHead {
		t.Errorf("head = %q, want %q", head, wantHead)
	}
	if body != wantBody {
		t.Errorf("body = %q, want %q", body, wantBody)
	}
}
