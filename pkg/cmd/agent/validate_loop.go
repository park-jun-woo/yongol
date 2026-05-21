//ff:func feature=agent type=command control=iteration dimension=1
//ff:what validateLoop — v2 validate 루프 (scaffold 후 에러 수렴)

package agent

import (
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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

	// Track per-file consecutive identical diagnostics for stall detection.
	// Key: relative filename, Value: {lastMessages, count}.
	stalled := map[string]*stallTracker{}

	for round := 1; round <= maxRounds; round++ {
		// 1. Validate.
		diags, err := runValidate(specsDir)
		if err != nil {
			return fmt.Errorf("round %d: %w", round, err)
		}

		allErrors := countErrors(diags)
		immutableCount := countImmutable(diags)
		fixable := filterImmutable(diags)
		fixableErrors := countErrors(fixable)

		fmt.Fprintf(out, "\nyongol agent: round %d — %d errors (%d fixable, %d immutable)\n",
			round, allErrors, fixableErrors, immutableCount)

		// Log immutable skips.
		logImmutableSkips(out, diags, absSpecs)

		if fixableErrors == 0 {
			elapsed := time.Since(start)
			fmt.Fprintf(out, "  scaffold + %d rounds (%s)\n", round, formatDuration(elapsed))
			return nil
		}

		// 2. Group fixable diagnostics by file.
		groups := groupByFile(fixable, absSpecs)
		sortByLayerPriority(groups)

		roundFixed := 0
		for _, g := range groups {
			// Check stall.
			key := g.relFile
			msgs := diagMessages(g.diags)
			joined := strings.Join(msgs, "\n")

			tracker, ok := stalled[key]
			if !ok {
				tracker = &stallTracker{}
				stalled[key] = tracker
			}
			if tracker.lastMessages == joined {
				tracker.count++
			} else {
				tracker.lastMessages = joined
				tracker.count = 1
			}
			if tracker.count >= 3 {
				fmt.Fprintf(out, "  stalled: %s (same errors 3 rounds — skipping)\n", g.relFile)
				continue
			}

			// 3. Fix.
			if err := fixFile(specsDir, ff, g.relFile, g.diags, llmFn, cfg); err != nil {
				fmt.Fprintf(out, "  skipped: %s (%v)\n", g.relFile, err)
				continue
			}

			hint := g.relFile
			if len(g.diags) > 0 {
				ruleID := extractRuleID(g.diags[0].Message)
				if ruleID != "" {
					hint += " (" + ruleID + ")"
				}
			}
			fmt.Fprintf(out, "  fixed: %s\n", hint)
			roundFixed++
			totalFixed++
		}

		if roundFixed == 0 {
			fmt.Fprintf(out, "  no files fixed this round — stopping early\n")
			break
		}
	}

	// Final validation to report remaining errors.
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

// stallTracker tracks consecutive identical diagnostics for a file.
type stallTracker struct {
	lastMessages string
	count        int
}

// logImmutableSkips prints a single summary line per immutable file.
func logImmutableSkips(w io.Writer, diags []diagnostic.Diagnostic, absSpecs string) {
	seen := map[string]bool{}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		if !isImmutable(d.File) {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		if seen[rel] {
			continue
		}
		seen[rel] = true
		fmt.Fprintf(w, "  skipped: %s (immutable)\n", rel)
	}
}

// collectRemainingFiles returns sorted unique file:line summaries for remaining errors.
func collectRemainingFiles(diags []diagnostic.Diagnostic, absSpecs string) []string {
	seen := map[string]bool{}
	var result []string
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		entry := fmt.Sprintf("%s:%d %s", rel, d.Line, d.Message)
		if seen[entry] {
			continue
		}
		seen[entry] = true
		result = append(result, entry)
	}
	sort.Strings(result)
	return result
}

// formatDuration returns a human-readable duration string.
func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%dm %02ds", m, s)
}
