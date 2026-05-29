//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what needsCurrentUser 단위 테스트 (시퀀스 중 currentUser 참조 존재 검사)

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestNeedsCurrentUser(t *testing.T) {
	cases := []struct {
		name string
		sf   ssacparser.ServiceFunc
		want bool
	}{
		{"no sequences", ssacparser.ServiceFunc{}, false},
		{"auth sequence present", ssacparser.ServiceFunc{
			Sequences: []ssacparser.Sequence{{Type: "get"}, {Type: "auth"}},
		}, true},
		{"currentUser input present", ssacparser.ServiceFunc{
			Sequences: []ssacparser.Sequence{
				{Type: "post", Inputs: map[string]string{"owner": "currentUser.ID"}},
			},
		}, true},
		{"none reference currentUser", ssacparser.ServiceFunc{
			Sequences: []ssacparser.Sequence{
				{Type: "get"},
				{Type: "post", Inputs: map[string]string{"title": "request.title"}},
			},
		}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := needsCurrentUser(tc.sf); got != tc.want {
				t.Errorf("needsCurrentUser = %v, want %v", got, tc.want)
			}
		})
	}
}
