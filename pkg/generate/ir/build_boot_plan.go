//ff:func feature=gen-ir type=generator control=sequence
//ff:what BuildBootPlan -- manifest + prepared.State → BootPlan 변환 (25블록 활성화 정규화)

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// BuildBootPlan inspects manifest + prepared.State to determine which of
// the 25 bootstrap blocks are active and assembles a BootPlan. The four
// activation patterns (A: Active predicate, B: prepared.ActiveBackends,
// C: internal conditional, D: always active) are collapsed into a single
// BootBlock.Active bool so downstream renderers never re-evaluate
// manifest conditions.
//
// Config fields are nil for now; they will be populated when renderers
// need typed configuration (Phase008+).
func BuildBootPlan(fs *yongol.Fullstack, ps *prepared.State, backendType string) *BootPlan {
	plan := &BootPlan{
		ProjectID:   projectID(fs),
		Module:      modulePath(fs),
		BackendType: backendType,
		ErrorEnvelope: &ErrorEnvelopeSpec{
			Fields: []string{"error", "message", "request_id"},
		},
	}

	plan.ActiveBlocks = buildActiveBlocks(fs, ps)
	return plan
}
