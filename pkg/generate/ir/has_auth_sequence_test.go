//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasAuthSequence(t *testing.T) {
	withAuth := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "auth"}}},
	}}
	if !hasAuthSequence(withAuth) {
		t.Errorf("expected auth sequence found")
	}
	without := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Type: "get"}}},
	}}
	if hasAuthSequence(without) {
		t.Errorf("expected no auth sequence")
	}
}
