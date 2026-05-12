//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-statemachine
//ff:what XSM-71 — @state input value type must be string-compatible (statemachine params are string)

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm71StateInputType validates XSM-71: every @state input value must
// resolve to a string-compatible Go type. Statemachine functions accept
// status parameters as string; passing pgtype.UUID or int64 causes a build
// failure that validate should catch early.
func xsm71StateInputType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			diags = append(diags, checkStateInputType(g, fn, seq)...)
		}
	}
	return diags
}
