//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what sequenceUsesCurrentUser 단위 테스트 (@auth 또는 currentUser. 입력 참조 검사)

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestSequenceUsesCurrentUser(t *testing.T) {
	cases := []struct {
		name string
		seq  ssacparser.Sequence
		want bool
	}{
		{"auth always uses currentUser", ssacparser.Sequence{Type: "auth"}, true},
		{"input prefixed with currentUser.", ssacparser.Sequence{
			Type:   "post",
			Inputs: map[string]string{"owner": "currentUser.ID"},
		}, true},
		{"plain input no currentUser", ssacparser.Sequence{
			Type:   "post",
			Inputs: map[string]string{"title": "request.title"},
		}, false},
		{"no inputs", ssacparser.Sequence{Type: "get"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sequenceUsesCurrentUser(tc.seq); got != tc.want {
				t.Errorf("sequenceUsesCurrentUser = %v, want %v", got, tc.want)
			}
		})
	}
}
