//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what validateLoopRound — validate 루프 1회차 실행 (진단 → 그룹핑 → 수정)

package agent

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func validateLoopRound(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer, absSpecs string, round int, stalled map[string]*stallTracker, totalFixed *int) (validateLoopRoundResult, error) {
	diags, err := runValidate(specsDir)
	if err != nil {
		return validateLoopRoundResult{}, fmt.Errorf("round %d: %w", round, err)
	}

	allErrors := countErrors(diags)
	immutableCount := countImmutable(diags)
	fixable := filterImmutable(diags)
	fixableErrors := countErrors(fixable)

	fmt.Fprintf(out, "\nyongol agent: round %d — %d errors (%d fixable, %d immutable)\n",
		round, allErrors, fixableErrors, immutableCount)

	logImmutableSkips(out, diags, absSpecs)

	if fixableErrors == 0 {
		return validateLoopRoundResult{allErrors: allErrors, fixable: 0, done: true}, nil
	}

	groups := groupByFile(fixable, absSpecs)
	sortByLayerPriority(groups)

	roundFixed := 0
	for _, g := range groups {
		if fixFileGroup(specsDir, ff, g, llmFn, cfg, out, stalled) {
			roundFixed++
			*totalFixed++
		}
	}

	return validateLoopRoundResult{allErrors: allErrors, fixable: fixableErrors, roundFixed: roundFixed}, nil
}
