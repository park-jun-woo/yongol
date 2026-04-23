//ff:func feature=orchestrator type=loader control=iteration dimension=1
//ff:what detectDirSSOT — 단일 dirSSOT 의 glob match 수를 세어 Presence 를 결정
package yongol

import (
	"fmt"
	"path/filepath"
)

// detectDirSSOT counts filesystem matches across the candidate's glob
// patterns and reports the resulting DetectedSSOT. Glob errors (effectively
// filepath.ErrBadPattern only) are surfaced as a hard error so a future
// refactor that introduces a caller-supplied pattern cannot silently pass.
func detectDirSSOT(d dirSSOT) (DetectedSSOT, error) {
	count := 0
	for _, g := range d.globs {
		pattern := filepath.Join(d.dir, g)
		matches, err := filepath.Glob(pattern)
		if err != nil {
			// filepath.Glob only returns ErrBadPattern (syntax error).
			// Patterns are hard-coded so this is effectively unreachable,
			// but surface it as a diagnostic to prevent silent pass.
			return DetectedSSOT{}, fmt.Errorf("detect SSOTs glob failed for %s: %w", pattern, err)
		}
		count += len(matches)
	}
	p := dirPresence(d.dir, count)
	if p == SSOTAbsent {
		return DetectedSSOT{Presence: SSOTAbsent}, nil
	}
	return DetectedSSOT{Kind: d.kind, Path: d.dir, Presence: p}, nil
}
