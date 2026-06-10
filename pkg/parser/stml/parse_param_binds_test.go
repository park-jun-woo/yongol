//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseParamBinds — 공유 코어의 분할·트림·소스 검사 위임·세그먼트 생략 검증

package stml

import (
	"fmt"
	"testing"
)

func TestParseParamBinds(t *testing.T) {
	allow := func(string) error { return nil }

	// Pairs are split on commas, both sides trimmed.
	binds, err := parseParamBinds(" a -> X ,  b ", "test param", allow)
	if err != nil {
		t.Fatalf("ok: unexpected error: %v", err)
	}
	if len(binds) != 2 || binds[0] != (LinkParamBind{Source: "a", Segment: "X"}) || binds[1] != (LinkParamBind{Source: "b"}) {
		t.Errorf("ok: got %+v", binds)
	}

	// The source check is delegated and its error propagates verbatim.
	deny := func(source string) error { return fmt.Errorf("denied %q", source) }
	if _, err := parseParamBinds("a -> X", "test param", deny); err == nil || err.Error() != `denied "a"` {
		t.Errorf("deny: got %v", err)
	}

	// Empty binding, empty segment after the arrow, double arrow.
	if _, err := parseParamBinds("", "test param", allow); err == nil {
		t.Errorf("empty: expected error, got nil")
	}
	if _, err := parseParamBinds("a ->", "test param", allow); err == nil {
		t.Errorf("empty segment: expected error, got nil")
	}
	if _, err := parseParamBinds("a -> X -> Y", "test param", allow); err == nil {
		t.Errorf("double arrow: expected error, got nil")
	}
}
