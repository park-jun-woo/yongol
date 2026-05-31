//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestScanSSaCAuthLocations(t *testing.T) {
	locs := map[[2]string]PairLocation{}
	fn := ssac.ServiceFunc{
		FileName: "svc.ssac",
		Sequences: []ssac.Sequence{
			authSeq("delete", "project", 3),
			authSeq("delete", "project", 9), // dup
		},
	}
	scanSSaCAuthLocations(fn, locs)
	loc := locs[[2]string{"delete", "project"}]
	if loc.File != "svc.ssac" || loc.Line != 3 {
		t.Errorf("loc = %+v, want svc.ssac:3", loc)
	}
}
