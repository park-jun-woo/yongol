//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseCapture — data-capture 값 파싱의 정상·구문 위반·sink 화이트리스트 검증

package stml

import "testing"

func TestParseCapture(t *testing.T) {
	// Single binding.
	binds, err := ParseCapture("access_token -> auth.token")
	if err != nil {
		t.Fatalf("single: unexpected error: %v", err)
	}
	if len(binds) != 1 || binds[0].RespField != "access_token" || binds[0].Sink != "auth.token" {
		t.Errorf("single: got %+v", binds)
	}

	// Multiple bindings with refresh sink.
	binds, err = ParseCapture("access_token -> auth.token, refresh_token -> auth.refresh")
	if err != nil {
		t.Fatalf("multi: unexpected error: %v", err)
	}
	if len(binds) != 2 || binds[1].RespField != "refresh_token" || binds[1].Sink != "auth.refresh" {
		t.Errorf("multi: got %+v", binds)
	}

	// Disallowed sink (session.* collides with the SSaC built-in).
	if _, err := ParseCapture("access_token -> session.token"); err == nil {
		t.Errorf("session sink: expected error, got nil")
	}

	// Missing arrow.
	if _, err := ParseCapture("access_token auth.token"); err == nil {
		t.Errorf("missing arrow: expected error, got nil")
	}

	// Empty response field.
	if _, err := ParseCapture("-> auth.token"); err == nil {
		t.Errorf("empty field: expected error, got nil")
	}

	// Trailing comma leaves an empty binding.
	if _, err := ParseCapture("access_token -> auth.token,"); err == nil {
		t.Errorf("trailing comma: expected error, got nil")
	}

	// Empty value.
	if _, err := ParseCapture(""); err == nil {
		t.Errorf("empty value: expected error, got nil")
	}
}
