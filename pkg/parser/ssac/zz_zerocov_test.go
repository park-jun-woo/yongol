//ff:func feature=ssac-parse type=test
//ff:what zz_zerocov_test — parseCallExpr / parseVerifyPassword 0% 커버리지 단위 테스트
package ssac

import (
	"strings"
	"testing"
)

func TestParseCallExpr_ZeroCov(t *testing.T) {
	// no paren → whole string is model, no args.
	model, args := parseCallExpr("Course.FindByID")
	if model != "Course.FindByID" || args != nil {
		t.Errorf("no-paren: model=%q args=%v", model, args)
	}

	// empty arg list → model only.
	model, args = parseCallExpr("session.Touch()")
	if model != "session.Touch" || args != nil {
		t.Errorf("empty-args: model=%q args=%v", model, args)
	}

	// with args.
	model, args = parseCallExpr("Course.FindByID(request.CourseID)")
	if model != "Course.FindByID" {
		t.Errorf("with-args: model=%q", model)
	}
	if len(args) != 1 {
		t.Fatalf("with-args: expected 1 arg, got %d", len(args))
	}
}

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

func TestParseVerifyPassword_ZeroCov(t *testing.T) {
	const ok = `User.email=request.body.email User.password_hash vs request.body.password -> user 401 "Invalid credentials"`
	seq, err := parseVerifyPassword(ok)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seq.Type != SeqVerifyPassword {
		t.Errorf("type=%q", seq.Type)
	}
	if seq.Model != "User" || seq.EmailCol != "email" || seq.HashCol != "password_hash" {
		t.Errorf("model/cols = %q/%q/%q", seq.Model, seq.EmailCol, seq.HashCol)
	}
	if seq.EmailExpr != "request.body.email" || seq.PasswordExpr != "request.body.password" {
		t.Errorf("exprs = %q / %q", seq.EmailExpr, seq.PasswordExpr)
	}
	if seq.ErrStatus != 401 || seq.Message != "Invalid credentials" {
		t.Errorf("status/msg = %d / %q", seq.ErrStatus, seq.Message)
	}
	if seq.Result == nil || seq.Result.Var != "user" || seq.Result.Type != "User" {
		t.Errorf("result = %+v", seq.Result)
	}

	// Error branches.
	errCases := map[string]string{
		"no arrow":         `User.email=x User.h vs p`,
		"no vs":            `User.email=x User.h -> u 401 "m"`,
		"empty pw":         `User.email=x User.h vs   -> u 401 "m"`,
		"one clause":       `User.email=x vs p -> u 401 "m"`,
		"email no eq":      `User.email User.h vs p -> u 401 "m"`,
		"email key 1part":  `email=x User.h vs p -> u 401 "m"`,
		"hash 1part":       `User.email=x hash vs p -> u 401 "m"`,
		"model mismatch":   `User.email=x Other.h vs p -> u 401 "m"`,
		"rhs too short":    `User.email=x User.h vs p -> u`,
		"status not int":   `User.email=x User.h vs p -> u abc "m"`,
	}
	for name, in := range errCases {
		if _, err := parseVerifyPassword(in); err == nil {
			t.Errorf("%s: expected error", name)
		} else if !strings.Contains(err.Error(), "verify-password") {
			t.Errorf("%s: unexpected error message %q", name, err)
		}
	}
}
