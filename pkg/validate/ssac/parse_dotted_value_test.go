//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what parseDottedValue — "source.field" 분리 검증 (정상/도트 없음/빈 문자열/다중 도트)

package ssac

import "testing"

func TestParseDottedValue(t *testing.T) {
	cases := []struct {
		name       string
		in         string
		wantSource string
		wantField  string
	}{
		{name: "normal", in: "course.InstructorID", wantSource: "course", wantField: "InstructorID"},
		{name: "no dot", in: "course", wantSource: "", wantField: ""},
		{name: "empty", in: "", wantSource: "", wantField: ""},
		{name: "dot at start", in: ".field", wantSource: "", wantField: "field"},
		{name: "dot at end", in: "source.", wantSource: "source", wantField: ""},
		{name: "multiple dots first wins", in: "a.b.c", wantSource: "a", wantField: "b.c"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			src, fld := parseDottedValue(c.in)
			if src != c.wantSource || fld != c.wantField {
				t.Errorf("parseDottedValue(%q) = (%q, %q), want (%q, %q)",
					c.in, src, fld, c.wantSource, c.wantField)
			}
		})
	}
}
