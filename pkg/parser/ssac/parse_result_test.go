//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseResult/parseGuard/parseCRUD(no-result) 파싱 검증
package ssac

import (
	"testing"
)

func TestParseResult(t *testing.T) {
	t.Run("plain type var", func(t *testing.T) {
		r := parseResult("Course course")
		if r == nil || r.Type != "Course" || r.Var != "course" || r.Wrapper != "" {
			t.Errorf("r = %+v", r)
		}
	})
	t.Run("wrapper generic", func(t *testing.T) {
		r := parseResult("Page[Gig] gigs")
		if r == nil || r.Wrapper != "Page" || r.Type != "Gig" || r.Var != "gigs" {
			t.Errorf("r = %+v", r)
		}
	})
	t.Run("slice type", func(t *testing.T) {
		r := parseResult("[]Course courses")
		if r == nil || r.Type != "[]Course" || r.Var != "courses" {
			t.Errorf("r = %+v", r)
		}
	})
	t.Run("malformed returns nil", func(t *testing.T) {
		if r := parseResult("only-one-token"); r != nil {
			t.Errorf("expected nil, got %+v", r)
		}
		if r := parseResult("a b c"); r != nil {
			t.Errorf("expected nil for 3 tokens, got %+v", r)
		}
	})
}
