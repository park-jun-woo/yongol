//ff:func feature=agent type=command control=sequence
//ff:what Run — agent 메인 흐름 (scaffold → v2 validate loop)

package agent

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// Run executes the agent: scaffold SSOT files, then run the v2 validate loop.
func Run(w io.Writer, cfg Config) error {
	start := time.Now()

	// Scaffold: generate SSOT files from features.yaml before validation
	ff, _ := features.Load(cfg.SpecsDir)
	if ff != nil {
		if err := scaffold(cfg.SpecsDir, ff, llmCall, cfg, w); err != nil {
			return fmt.Errorf("scaffold: %w", err)
		}
	}
	if cfg.MaxRounds == 0 {
		fmt.Fprintf(w, "\nyongol agent: scaffold only (max-rounds=0), %.1fs\n", time.Since(start).Seconds())
		return nil
	}

	// v2 validate loop: validate → filter immutable → fix per file → repeat.
	return validateLoop(cfg.SpecsDir, ff, llmCall, cfg, w, os.Stderr, cfg.MaxRounds)
}
