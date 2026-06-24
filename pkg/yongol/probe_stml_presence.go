//ff:func feature=orchestrator type=loader control=sequence dimension=1
//ff:what 도메인 frontend 디렉토리의 *.html glob 으로 STML Presence 3-상태 판정
package yongol

// probeSTMLPresence reports the three-state presence of a domain's STML
// frontend directory, reusing detectDirSSOT's glob-count classification with the
// same "*.html" pattern as the shared STML dirSSOT (directory_ssots.go). Glob
// errors are structurally unreachable (hard-coded "*.html" pattern), so the
// returned error is intentionally discarded — detectDirSSOT yields a zero-value
// SSOTAbsent Presence on the impossible error path anyway.
func probeSTMLPresence(dir string) SSOTPresence {
	entry, _ := detectDirSSOT(dirSSOT{kind: KindSTML, dir: dir, globs: []string{"*.html"}})
	return entry.Presence
}
