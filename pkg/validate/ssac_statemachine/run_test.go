//ff:func feature=validate type=test control=sequence topic=ssac-statemachine
//ff:what TestRun — Run SSaC↔StateMachine 교차 검증 집계 호출 검증

package ssac_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: &openapi3.T{
			OpenAPI: "3.0.0",
			Info:    &openapi3.Info{Title: "t", Version: "1"},
			Paths:   openapi3.NewPaths(),
		},
	}
	fs.SetGround(&rule.Ground{
		Lookup: map[string]rule.StringSet{},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	})
	// Empty fixture: aggregator should run all sub-rules without panicking.
	diags := Run(fs)
	if diags == nil {
		// nil is acceptable (no diagnostics); just ensure no panic occurred.
		return
	}
}
