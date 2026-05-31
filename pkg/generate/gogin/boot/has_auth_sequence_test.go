//ff:func feature=gen-gogin type=test control=sequence
//ff:what hasAuthSequence — SSaC ServiceFuncs 에 @auth 시퀀스 존재 여부
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasAuthSequence(t *testing.T) {
	none := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Name: "GetX", Sequences: []ssac.Sequence{{Type: "get"}, {Type: "response"}}},
	}}
	if hasAuthSequence(none) {
		t.Errorf("no @auth sequence should be false")
	}

	with := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Name: "GetX", Sequences: []ssac.Sequence{{Type: "get"}}},
		{Name: "Login", Sequences: []ssac.Sequence{{Type: "auth"}, {Type: "response"}}},
	}}
	if !hasAuthSequence(with) {
		t.Errorf("an @auth sequence should be true")
	}

	if hasAuthSequence(&yongol.Fullstack{}) {
		t.Errorf("no funcs should be false")
	}
}
