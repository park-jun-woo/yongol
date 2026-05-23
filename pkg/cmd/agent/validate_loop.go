//ff:func feature=agent type=command control=iteration dimension=1
//ff:what validateLoop — v2 validate 루프 (scaffold 후 에러 수렴)

package agent

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// validateLoop runs the v2 validate→fix cycle up to maxRounds times.
// Each round: validate → filter immutable → group by file → fix each file via LLM.
// Stops early when 0 fixable errors remain or no progress is made.
func validateLoop(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out, errOut io.Writer, maxRounds int) error {
	start := time.Now()
	absSpecs, err := filepath.Abs(specsDir)
	if err != nil {
		return fmt.Errorf("resolve specs dir: %w", err)
	}

	totalFixed := 0
	stalled := map[string]*stallTracker{}

	for round := 1; round <= maxRounds; round++ {
		result, err := validateLoopRound(specsDir, ff, llmFn, cfg, out, absSpecs, round, stalled, &totalFixed)
		if err != nil {
			return err
		}
		if result.done {
			elapsed := time.Since(start)
			fmt.Fprintf(out, "  scaffold + %d rounds (%s)\n", round, formatDuration(elapsed))
			return nil
		}
		if result.roundFixed == 0 {
			fmt.Fprintf(out, "  no files fixed this round — stopping early\n")
			break
		}
	}

	diags, _ := runValidate(specsDir)
	fixable := filterImmutable(diags)
	remaining := countErrors(fixable)
	elapsed := time.Since(start)

	if remaining > 0 {
		fmt.Fprintf(out, "\nyongol agent: %d errors remaining after %d rounds (%s)\n",
			remaining, maxRounds, formatDuration(elapsed))
		remainingFiles := collectRemainingFiles(fixable, absSpecs)
		for _, f := range remainingFiles {
			fmt.Fprintf(out, "  %s\n", f)
		}
		return fmt.Errorf("%d errors remaining", remaining)
	}

	fmt.Fprintf(out, "\nyongol agent: 0 errors — %d files fixed (%s)\n", totalFixed, formatDuration(elapsed))
	return nil
}
