//ff:func feature=validate type=command control=iteration dimension=1
//ff:what Validate — Fullstack 기반 per-SSOT + pair 검증 실행, Report 반환
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/ground"
	contractvalidate "github.com/park-jun-woo/yongol/pkg/validate/contract"
)

// Validate runs per-SSOT and pair-cross validation and returns a Report.
// Parser failures surface as nil fields in fs; corresponding steps emit
// StatusMissing.
//
// When WithArtsDir(dir) is supplied and the directory exists, the
// contract validator (PRV-01 / PRV-02) is appended as a final step so
// preserve-drift is surfaced alongside SSOT drift.
func Validate(fs *yongol.Fullstack, opts ...Option) *Report {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.ground == nil {
		cfg.ground = ground.Build(fs)
	}
	fs.SetGround(cfg.ground)
	r := &Report{}
	for _, s := range allSteps() {
		r.Steps = append(r.Steps, s.execute(fs, cfg))
	}
	if cfg.artsDir != "" {
		diags := contractvalidate.Run(fs, cfg.artsDir)
		r.Steps = append(r.Steps, finalize("contract", diags))
	}
	return r
}
