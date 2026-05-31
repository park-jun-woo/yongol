//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증
package yongol

func endsWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
