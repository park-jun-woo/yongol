//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestFirstSSaCAuthLocation(t *testing.T) {
	funcs := []ssac.ServiceFunc{{FileName: "s.ssac", Sequences: []ssac.Sequence{authSeq("write", "doc", 4)}}}
	locs := firstSSaCAuthLocation(funcs)
	if locs[[2]string{"write", "doc"}].Line != 4 {
		t.Errorf("locs = %+v", locs)
	}
}
