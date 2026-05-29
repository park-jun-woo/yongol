//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what needsTransaction 단위 테스트 (mutating seq 존재 시 true)

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestNeedsTransaction(t *testing.T) {
	cases := []struct {
		name  string
		types []string
		want  bool
	}{
		{"read-only get", []string{"get"}, false},
		{"post mutates", []string{"get", "post"}, true},
		{"put mutates", []string{"put"}, true},
		{"delete mutates", []string{"delete"}, true},
		{"guards only", []string{"empty", "exists", "auth", "response"}, false},
		{"empty sequences", nil, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var seqs []ssacparser.Sequence
			for _, ty := range tc.types {
				seqs = append(seqs, ssacparser.Sequence{Type: ty})
			}
			sf := ssacparser.ServiceFunc{Sequences: seqs}
			if got := needsTransaction(sf); got != tc.want {
				t.Errorf("needsTransaction(%v) = %v, want %v", tc.types, got, tc.want)
			}
		})
	}
}
