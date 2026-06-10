//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-10 — 모든 rate_limit 항목은 Rate≥1 + Period 파싱 가능 (제로값/빈 항목 차단)

package manifest

import (
	"fmt"
	"sort"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c10RateLimitValueValid enforces that every backend.rate_limit entry carries
// usable values: Rate must be >= 1 and Period must parse via time.ParseDuration.
// C-7 (section present) and C-8 (Login key present) only check existence, so a
// zero-value entry ({Rate:0, Period:""}) — e.g. one that slipped through a loose
// YAML decode — would otherwise pass validate and then be silently dropped by
// codegen (block_rate_limit.go continues on ParseDuration failure), producing a
// backend with no rate limiter. Flagging these at validate time blocks input
// that codegen would discard without warning.
func c10RateLimitValueValid(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if len(fs.Manifest.Backend.RateLimit) == 0 {
		// C-7 already flags a missing/empty rate_limit section.
		return nil
	}

	// Iterate in a deterministic order so diagnostics are stable.
	opIDs := make([]string, 0, len(fs.Manifest.Backend.RateLimit))
	for opID := range fs.Manifest.Backend.RateLimit {
		opIDs = append(opIDs, opID)
	}
	sort.Strings(opIDs)

	var diags []diagnostic.Diagnostic
	for _, opID := range opIDs {
		entry := fs.Manifest.Backend.RateLimit[opID]
		if entry.Rate < 1 {
			diags = append(diags, diagnostic.Diagnostic{
				File:  "manifest.yaml",
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[C-10] backend.rate_limit.%s has rate=%d — rate must be >= 1",
					opID, entry.Rate),
				Advice: "Set a positive rate (requests allowed per period). " +
					"A zero rate disables the limiter and is silently dropped by " +
					"codegen. Example:\n" +
					"  backend:\n" +
					"    rate_limit:\n" +
					"      " + opID + ":\n" +
					"        rate: 5\n" +
					"        period: \"1m\"\n" +
					"        key: ip",
			})
		}
		if _, err := time.ParseDuration(entry.Period); err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:  "manifest.yaml",
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[C-10] backend.rate_limit.%s has period=%q — period must be a "+
						"valid Go duration (time.ParseDuration)",
					opID, entry.Period),
				Advice: "Set a parseable duration such as \"1m\", \"30s\", or \"1h\". " +
					"An empty or invalid period is silently dropped by codegen, " +
					"leaving the endpoint unprotected. Example:\n" +
					"  backend:\n" +
					"    rate_limit:\n" +
					"      " + opID + ":\n" +
					"        rate: 5\n" +
					"        period: \"1m\"\n" +
					"        key: ip",
			})
		}
	}
	return diags
}
