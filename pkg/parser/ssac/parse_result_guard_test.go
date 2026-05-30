//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what parseResult/parseGuard/parseCRUD(no-result) 파싱 검증

package ssac

import "testing"

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

func TestParseGuard(t *testing.T) {
	t.Run("target and message", func(t *testing.T) {
		seq := parseGuard("empty", `course "not found"`)
		if seq.Type != "empty" || seq.Target != "course" || seq.Message != "not found" {
			t.Errorf("seq = %+v", seq)
		}
		if seq.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", seq.ErrStatus)
		}
	})
	t.Run("with explicit status remainder", func(t *testing.T) {
		seq := parseGuard("exists", `dup "conflict" 422`)
		if seq.ErrStatus != 422 {
			t.Errorf("ErrStatus = %d, want 422", seq.ErrStatus)
		}
	})
	t.Run("non-numeric remainder ignored", func(t *testing.T) {
		seq := parseGuard("empty", `x "msg" notanumber`)
		if seq.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", seq.ErrStatus)
		}
	})
}

func TestParseCRUD_NoResult(t *testing.T) {
	seq, err := parseCRUD("delete", "Reservation.Delete({ID: request.ID})", false)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if seq.Type != "delete" || seq.Model != "Reservation.Delete" {
		t.Errorf("seq = %+v", seq)
	}
	if seq.Inputs["ID"] != "request.ID" {
		t.Errorf("inputs = %v", seq.Inputs)
	}
}
