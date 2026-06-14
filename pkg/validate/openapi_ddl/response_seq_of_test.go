//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what responseSeqOf — Sequences 에서 @response 시퀀스 탐색, 없으면 nil

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResponseSeqOf(t *testing.T) {
	fn := &ssac.ServiceFunc{Sequences: []ssac.Sequence{
		{Type: "get"},
		{Type: "response", Target: "rule"},
	}}
	if seq := responseSeqOf(fn); seq == nil || seq.Target != "rule" {
		t.Errorf("expected response seq with Target rule, got %+v", seq)
	}
	noResp := &ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}}}
	if seq := responseSeqOf(noResp); seq != nil {
		t.Errorf("expected nil, got %+v", seq)
	}
}
