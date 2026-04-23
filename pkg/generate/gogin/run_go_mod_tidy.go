//ff:func feature=gen-gogin type=util control=sequence
//ff:what runGoModTidy — `go mod tidy` 실행 (stderr 캡처는 runGo 가 담당)

package gogin

// runGoModTidy executes `go mod tidy` in dir and captures stderr. On failure
// it returns a wrapped error carrying the (truncated) stderr so callers can
// surface the concrete cause (network error, auth failure, unresolved
// module, etc.) instead of silently producing a broken go.mod.
func runGoModTidy(dir string) error {
	return runGo(dir, "mod", "tidy")
}
