//ff:func feature=chain type=test control=sequence
//ff:what TestTracePolicyNoAuth — @auth resource 가 없으면 nil 반환
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTracePolicyNoAuth(t *testing.T) {
	specsDir, regoFile := tracePolicySetup(t)
	policies := []rego.Policy{{File: regoFile, Rules: []rego.AllowRule{{Resource: "project", Actions: []string{"delete"}}}}}
	sfNone := &ssac.ServiceFunc{Name: "X", Sequences: []ssac.Sequence{{Type: "get", Model: "Y.Z"}}}
	if tracePolicy(sfNone, policies, specsDir) != nil {
		t.Error("expected nil when no @auth resources")
	}
}
