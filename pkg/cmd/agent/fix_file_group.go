//ff:func feature=agent type=helper control=sequence
//ff:what fixFileGroup — 파일 그룹의 stall 확인 후 fixFile 호출

package agent

import (
	"fmt"
	"io"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func fixFileGroup(specsDir string, ff *features.FeaturesFile, g fileGroup, llmFn LLMCallFunc, cfg Config, out io.Writer, stalled map[string]*stallTracker) bool {
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
		return false
	}

	if err := fixFile(specsDir, ff, g.relFile, g.diags, llmFn, cfg); err != nil {
		fmt.Fprintf(out, "  skipped: %s (%v)\n", g.relFile, err)
		return false
	}

	hint := g.relFile
	if len(g.diags) > 0 {
		ruleID := extractRuleID(g.diags[0].Message)
		if ruleID != "" {
			hint += " (" + ruleID + ")"
		}
	}
	fmt.Fprintf(out, "  fixed: %s\n", hint)
	return true
}
