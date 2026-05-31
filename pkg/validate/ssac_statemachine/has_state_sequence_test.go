//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestHasStateSequence — hasStateSequence @state 시퀀스 존재 여부 검증
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestHasStateSequence(t *testing.T) {
	cases := []struct {
		name string
		seqs []ssac.Sequence
		want bool
	}{
		{name: "has state", seqs: []ssac.Sequence{{Type: "get"}, {Type: "state"}}, want: true},
		{name: "no state", seqs: []ssac.Sequence{{Type: "get"}, {Type: "post"}}, want: false},
		{name: "empty", seqs: nil, want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasStateSequence(tc.seqs); got != tc.want {
				t.Errorf("hasStateSequence() = %v, want %v", got, tc.want)
			}
		})
	}
}
