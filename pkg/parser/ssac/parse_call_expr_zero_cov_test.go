//ff:func feature=ssac-parse type=test control=sequence
//ff:what zz_zerocov_test — parseCallExpr / parseVerifyPassword 0% 커버리지 단위 테스트
package ssac

import (
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
