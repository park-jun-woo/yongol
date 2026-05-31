//ff:func feature=orchestrator type=util control=sequence
//ff:what findYongolPkgRoot — built-in FuncSpec(ssac/pkg) 디렉토리를 3단계 fallback 으로 해결

package yongol

import "os"

// findYongolPkgRoot resolves the directory that holds built-in FuncSpecs
// (auth, session, cache, file, mail, …). Priority:
//  1. `YONGOL_SSAC_PKG` env var (absolute path override for CI / container).
//  2. Sibling `ssac/pkg` of the yongol repo walked up from CWD (dev worktree).
//  3. `$GOMODCACHE/github.com/park-jun-woo/ssac@<version>/pkg` (go install env).
//
// Returns "" when none applies — callers treat that as "no built-ins loaded".
func findYongolPkgRoot() string {
	if p := os.Getenv("YONGOL_SSAC_PKG"); p != "" && isDir(p) {
		return p
	}
	if p := findYongolPkgRootFromCWD(); p != "" {
		return p
	}
	return findYongolPkgRootFromGoModCache()
}
