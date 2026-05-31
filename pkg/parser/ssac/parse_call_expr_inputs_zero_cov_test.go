//ff:func feature=ssac-parse type=test control=sequence
//ff:what zz_zerocov_test — parseCallExpr / parseVerifyPassword 0% 커버리지 단위 테스트
package ssac

import (
	"testing"
)

func TestParseCallExprInputs_ZeroCov(t *testing.T) {
	// no paren.
	m, in, rem, err := parseCallExprInputs("queue.Publish")
	if err != nil || m != "queue.Publish" || in != nil || rem != "" {
		t.Errorf("no-paren: %q %v %q %v", m, in, rem, err)
	}

	// no closing paren.
	m, in, rem, err = parseCallExprInputs("queue.Publish({a: b}")
	if err != nil || m != "queue.Publish" || in != nil || rem != "" {
		t.Errorf("no-close: %q %v %q %v", m, in, rem, err)
	}

	// empty inner with remainder.
	m, in, rem, err = parseCallExprInputs("queue.Publish() after")
	if err != nil || m != "queue.Publish" || in != nil || rem != "after" {
		t.Errorf("empty-inner: %q %v %q %v", m, in, rem, err)
	}

	// inputs present.
	m, in, rem, err = parseCallExprInputs("queue.Publish({delay: 1800}) tail")
	if err != nil {
		t.Fatalf("inputs: err=%v", err)
	}
	if m != "queue.Publish" || rem != "tail" || in["delay"] != "1800" {
		t.Errorf("inputs: %q in=%v rem=%q", m, in, rem)
	}

	// malformed inputs → error.
	if _, _, _, err := parseCallExprInputs("queue.Publish({bad})"); err == nil {
		t.Error("expected error for malformed inputs")
	}
}
