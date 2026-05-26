//ff:func feature=validate type=test control=iteration dimension=2 topic=ssac-structural
//ff:what s62unusedInSubsequent — 후속 시퀀스에서 변수 미사용 여부 검증 (Inputs/Fields/Target/start 이후만)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS62UnusedInSubsequent(t *testing.T) {
	seqs := []parsessac.Sequence{
		{Inputs: map[string]string{"id": "order.ID"}},
		{Fields: map[string]string{"name": "user.Name"}},
		{Target: "course.ID"},
		{Inputs: map[string]string{"status": "reservation.Status"}},
	}

	cases := []struct {
		name    string
		varName string
		start   int
		want    bool
	}{
		{name: "used in Inputs", varName: "order", start: 0, want: false},
		{name: "used in Fields", varName: "user", start: 0, want: false},
		{name: "used in Target", varName: "course", start: 0, want: false},
		{name: "not used", varName: "unknown", start: 0, want: true},
		{name: "used before start only", varName: "order", start: 1, want: true},
		{name: "start at last", varName: "reservation", start: 3, want: false},
		{name: "start past end", varName: "reservation", start: 4, want: true},
		{name: "empty seqs", varName: "x", start: 0, want: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := seqs
			if c.name == "empty seqs" {
				input = nil
			}
			got := s62unusedInSubsequent(c.varName, input, c.start)
			if got != c.want {
				t.Errorf("s62unusedInSubsequent(%q, seqs, %d) = %v, want %v",
					c.varName, c.start, got, c.want)
			}
		})
	}
}
